package indexer

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/dopedao/dope-monorepo/packages/api/indexer/processor"
	"github.com/dopedao/dope-monorepo/packages/api/internal/dbprovider"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/event"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/syncstate"
	"github.com/dopedao/dope-monorepo/packages/api/internal/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog/log"
	"github.com/withtally/synceth/engine"
)

const blockLimit = 500

// Limit for number of ent transactions to run concurrently
// If this is too high we run into deadlocks and the indexer crashes.
// This can also happen when code inside each "tx" function is slow.
const txLimit = 50

type Contract struct {
	Name       string
	Address    common.Address
	StartBlock uint64
	Processor  processor.Processor
}

type EthConfig struct {
	RPC       string
	Interval  time.Duration
	Contracts []Contract
}

type Ethereum struct {
	sync.Mutex
	latest    uint64
	ent       *ent.Client
	eth       engine.Client
	ticker    *time.Ticker
	contracts []*Contract
}

// Reads information from Ethereum and Optimism blockchains
// and attempts to persist information about DOPE, PAPER, HUSTLERS, and GEAR
func NewEthereumIndexer(ctx context.Context, entClient *ent.Client, config EthConfig) *Ethereum {
	ctx, log := logger.LogFor(ctx)

	retryableHTTPClient := retryablehttp.NewClient()
	retryableHTTPClient.Logger = nil
	c, err := rpc.DialHTTPWithClient(config.RPC, retryableHTTPClient.StandardClient())
	if err != nil {
		log.Fatal().Msg("Dialing ethereum rpc.") //nolint:gocritic
	}

	log.Info().Msgf("Starting client for %v", c)
	eth := ethclient.NewClient(c)

	e := &Ethereum{
		ent:    entClient,
		eth:    eth,
		ticker: time.NewTicker(config.Interval),
	}

	for _, c := range config.Contracts {
		c := c
		log.Info().Msgf("Get SyncState for %v", c.Name)
		s, err := entClient.SyncState.Get(ctx, c.Address.Hex())
		if err != nil && !ent.IsNotFound(err) {
			log.Fatal().Err(err).Msg("Fetching sync state.")
		} else if err == nil {
			c.StartBlock = s.StartBlock
		}

		if err := c.Processor.Setup(c.Address, eth); err != nil {
			log.Fatal().Msg("Setting up processor.")
		}
		e.contracts = append(e.contracts, &c)
	}

	return e
}

func eventLogCommitter(
	ctx context.Context,
	c *Contract,
	l types.Log,
	committer func(tx *ent.Tx) error,
) func(tx *ent.Tx) error {
	return func(tx *ent.Tx) error {
		id := fmt.Sprintf("%s-%s-%d", c.Address.Hex(), l.TxHash.Hex(), l.Index)
		if _, err := tx.Event.Get(ctx, id); err != nil {
			if ent.IsNotFound(err) {
				if err := tx.Event.Create().
					SetID(id).
					SetAddress(c.Address).
					SetHash(l.TxHash).
					SetIndex(uint64(l.Index)).
					OnConflictColumns(event.FieldID).
					UpdateNewValues().
					Exec(ctx); err != nil {
					return fmt.Errorf("%v - eventLogCommitter: creating event log %s: %w", c.Name, id, err)
				}

				return committer(tx)
			}

			return fmt.Errorf("%v - eventLogCommitter: getting event log %s: %w", c.Name, id, err)
		}
		log.Warn().Msgf("%v – eventLogCommitter: duplicate event log %s", c.Name, id)
		return committer(tx)
	}
}

func (e *Ethereum) Sync(ctx context.Context) {
	ctx, log := logger.LogFor(ctx)

	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			latest, err := e.eth.BlockNumber(ctx)
			if err != nil {
				log.Err(err).Msg("Getting latest block number.")
				continue
			}

			e.latest = latest
			numUpdates := 0
			var wg sync.WaitGroup
			for _, contract := range e.contracts {
				wg.Add(1)
				go func(contract *Contract) {
					defer wg.Done()
					syncErr := syncContract(ctx, contract, e, &numUpdates)
					if syncErr != nil {
						log.Fatal().Err(syncErr).Msgf("syncContract failed %s", contract.Name)
					}
				}(contract)
			}
			wg.Wait()
			e.Unlock()

			if numUpdates > 0 {
				if err := e.ent.RefreshSearchIndex(ctx); err != nil {
					log.Err(err).Msgf("Refreshing search index.")
				}
			}
			numUpdates = 0

		case <-ctx.Done():
			return
		}
	}
}

func syncContract(ctx context.Context, contract *Contract, e *Ethereum, numUpdates *int) error {
	_from := contract.StartBlock

	for {
		_to := Min(e.latest, _from+blockLimit)

		if _from > _to {
			// No new blocks
			break
		}

		log.Info().Msgf("Indexing %s from %d to %d.", contract.Name, _from, _to)

		logs, err := e.eth.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(_from),
			ToBlock:   new(big.Int).SetUint64(_to),
			Addresses: []common.Address{contract.Address},
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Filtering logs.")
		}

		var committers []func(*ent.Tx) error
		for _, l := range logs {
			committer, err := contract.Processor.ProcessElement(contract.Processor)(ctx, l)
			if err != nil {
				// Commented out so we don't crash the indexer on every mis-process
				log.Fatal().Err(err).Msgf("Processing element %s.", l.TxHash.Hex())
				// log.Err(err).Msgf(
				// 	"Contract %s error processing tx hash %s.", c.Address, l.TxHash.Hex())
				// break
			}

			if committer == nil {
				// Event log not handled
				continue
			}
			commitFunc := eventLogCommitter(ctx, contract, l, committer)
			committers = append(committers, commitFunc)
		}

		_from = _to + 1
		log.Debug().
			Msgf("%v - Processed %v", contract.Name, _to)

		// Because we can have quite a few transactions to commit, this code
		// can cause deadlocks. We chunk these into batches of 100 to avoid them.
		//
		// Because of this design, each "committer" should be idempotent
		// and ensure it uses an Ent upsert operation.
		txChunks := chunkTx(committers, txLimit)

		log.Debug().Msgf("%v - %v Chunks with %v total transactions", contract.Name, len(txChunks), len(committers))
		for _, committerChunk := range txChunks {
			if err := ent.WithTx(ctx, e.ent, func(tx *ent.Tx) error {
				log.Debug().Msgf(
					"%v - block %v Committing %v changes",
					contract.Name,
					_from,
					len(committerChunk))
				for _, c := range committerChunk {
					if err := c(tx); err != nil {
						log.Fatal().Err(err).Msgf("%s - FAIL committing transactions", contract.Name)
						return err
					}
				}
				log.Debug().Msgf(
					"%v - block %v COMMITTED %v changes",
					contract.Name,
					_from,
					len(committerChunk))
				*numUpdates += len(committerChunk)
				return nil
			}); err != nil {
				log.Err(err).Msgf("%s - FAIL Indexing contract", contract.Name)
				return err
			}
		}

		// Update sync state only when we've successfully processed
		// all the transactions above.
		log.Debug().Msgf("%v - Updating sync state %v", contract.Name, _from)
		if err := dbprovider.Ent().SyncState.
			Create().
			SetID(contract.Address.Hex()).
			SetStartBlock(_from).
			OnConflictColumns(syncstate.FieldID).
			UpdateStartBlock().
			Exec(ctx); err != nil {
			log.Err(err).Msgf("Updating sync state for %s", contract.Address.Hex())
			return err
		}

		if _to == e.latest {
			contract.StartBlock = _from
			break
		}
	}
	return nil
}

// Split slices into chunks of predetermined size
func chunkTx(slice []func(*ent.Tx) error, chunkSize int) [][]func(*ent.Tx) error {
	var chunks [][]func(*ent.Tx) error
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}
