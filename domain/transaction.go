package domain

import "math"

type Transaction struct {
	Price float64 `json:"price"`
	Date string `json:"date"`
	Quantity float64 `json:"quantity"`
	Asset string `json:"asset"`
	AssetPrice float64 `json:"asset_price"`
}

func (t Transaction) GetProfit(latestPrice float64) float64 {
	profit := ((latestPrice - t.AssetPrice) / t.AssetPrice) * 100

	return math.Round(profit*100)/100
}