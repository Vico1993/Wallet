package domain

type Transaction struct {
	Price float64 `json:"price"`
	Date string `json:"date"`
	Quantity float64 `json:"quantity"`
	Asset string `json:"asset"`
	AssetPrice float64 `json:"asset_price,omitempty"`
}

func (t Transaction) GetProfit(latestPrice float64) float64 {
	return calculProfit(t.GetAssetPrice(), latestPrice)
}

func (t Transaction) GetAssetPrice() float64 {
	if t.AssetPrice == 0 {
		return 1 * t.Price / t.Quantity
	}

	return t.AssetPrice
}