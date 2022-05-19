package service

import (
	"log"
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func GetAssetPrice(asset string) {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_API_KEY"),
	})

	// crypto, err := client.Cryptocurrency.Info(&cmc.InfoOptions{
	// 	Symbol: asset,
	// })

	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Convert: "CAD",
		Symbol: asset,
	})

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(quotes)

	for _, v := range quotes {
		log.Fatal(v.Quote["CAD"].Price)
	}

	// return price
}