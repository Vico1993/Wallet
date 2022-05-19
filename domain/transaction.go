package domain

type Transaction struct {
	Price float64 `json:"price"`
	Date string `json:"date"`
	Quantity float64 `json:"quantity"`
	Asset string `json:"asset"`
	AssetPrice float64 `json:"asset_price"`
}
