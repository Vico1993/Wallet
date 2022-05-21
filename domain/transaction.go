package domain

import "math"

type Transaction struct {
	Price float64 `json:"price"`
	Date string `json:"date"`
	Quantity float64 `json:"quantity"`
	Asset string `json:"asset"`
	AssetPrice float64 `json:"asset_price,omitempty"`
}

func (t Transaction) GetProfit(latestPrice float64) float64 {
	var assetPrice float64
	if assetPrice == 0 {
		assetPrice = t.GetAssetPrice()
	} else {
		assetPrice = t.AssetPrice
	}

	profit := ((latestPrice - assetPrice) / assetPrice) * 100

	return math.Round(profit*100)/100
}

func (t Transaction) GetAssetPrice() float64 {
	return 1 * t.Price / t.Quantity
}