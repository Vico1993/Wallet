package service

import (
	"errors"
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

var (
	prices = map[string]float64{}
)

func GetAssetPrice(asset string) (float64, error) {
	// To avoid hitting CMC during development
	if os.Getenv("DEBUG") == "1" {
		prices["BTC"] = 26325.07
		prices["ETH"] = 1447.83
		prices["VET"] = 0.02879
		prices["DOGE"] = 0.07984
		prices["BUSD"] = 1
		prices["MANA"] = 1.05
		prices["EGLD"] = 68.06
		prices["ERD"] = 0
	}

	if val, ok := prices[asset]; ok {
		return val, nil
	}

	if os.Getenv("CMC_API_KEY") == "" {
		return 0, errors.New("No CoinMarketCap API key found, please setup your .env")
	}

	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_API_KEY"),
	})

	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Convert: "CAD",
		Symbol: asset,
	})

	if err != nil {
		return 0, err
	}

	// If no price return
	if len(quotes) == 0 {
		return 0, errors.New("No price found for the Asset: " + asset)
	}

	prices[asset] = quotes[0].Quote["CAD"].Price

	return prices[asset], nil
}