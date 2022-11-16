package main

import (
	"context"
	"os"
	"sync"

	"github.com/dopedao/dope-monorepo/packages/api/indexer"
	"github.com/dopedao/dope-monorepo/packages/api/internal/dbprovider"
	"github.com/dopedao/dope-monorepo/packages/api/internal/envcfg"
	"github.com/rs/zerolog"
)

// Go Ethereum indexer for running from command line
// without the HTTP server wrapping necessary for deploying to GCP.
//
// At some point if we make a Cloud Run Docker container it should call
// this instead of the cmd/indexServer program
func main() {
	log := zerolog.New(os.Stderr)
	ctx := context.Background()
	var wg sync.WaitGroup
	for _, conf := range indexer.Config[envcfg.Network] {
		switch conf := conf.(type) {
		case indexer.EthConfig:
			log.Debug().Msgf("Starting %v", conf)

			eth := indexer.NewEthereumIndexer(
				log.WithContext(ctx),
				dbprovider.Ent(),
				conf)

			wg.Add(1)
			go func() {
				defer wg.Done()
				eth.Sync(log.WithContext(ctx))
			}()
		}
	}
	wg.Wait()
}
