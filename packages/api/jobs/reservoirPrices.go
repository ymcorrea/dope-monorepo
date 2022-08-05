package jobs

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/dopedao/dope-monorepo/packages/api/internal/dbprovider"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

type ReservoirTokenPrices struct {
	Tokens map[string]float64 `json:"tokens"`
}

func ReservoirPrices(queue chan int) {
	log := zerolog.New(os.Stderr)
	ctx := context.Background()

	contractAddress := "0x8707276df042e89669d69a177d3da7dc78bd8723"
	url := "https://api.reservoir.tools/tokens/floor/v1?collection=" + contractAddress

	resp, err := retryablehttp.Get(url)
	if err != nil {
		log.Fatal().Msgf("ReservoirPrices Get %+v.", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Msgf("ReservoirPrices ReadAll %+v.", err)
	}
	log.Info().Msg(string(body))

	var data ReservoirTokenPrices
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal().Msgf("Unmarshal %+v.", err)
	}

	// Update all to have default sale price
	dbprovider.Ent().
		Dope.
		Update().
		SetSalePrice(0).
		ExecX(ctx)

	// Loop through response and update.
	for tokenId, salePrice := range data.Tokens {
		log.Info().Msgf("Updating %s: %.2f", tokenId, salePrice)
		dbprovider.Ent().
			Dope.
			UpdateOneID(tokenId).
			SetSalePrice(salePrice).
			ExecX(ctx)
	}

	log.Info().Msg("DONE: ReservoirPrices")

	// Pop this job off the queue
	<-queue
}
