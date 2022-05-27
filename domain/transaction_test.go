package domain

import "testing"

func TestGetAssetPrice(t *testing.T) {
	table := []struct {
		input    Transaction
		expected float64
	}{
		{
			Transaction{
				Price: 100,
				Date: "2022-05-26",
				Quantity: 1,
				Asset: "BTC",
				AssetPrice: 100,
			},
			100,
		},
		{
			Transaction{
				Price: 100,
				Date: "2022-05-26",
				Quantity: 1,
				Asset: "BTC",
			},
			100,
		},
		{
			Transaction{
				Price: 100,
				Date: "2022-05-26",
				Quantity: 0.8,
				Asset: "BTC",
			},
			125,
		},
		{
			Transaction{
				Price: 100,
				Date: "2022-05-26",
				Quantity: 0.08,
				Asset: "BTC",
			},
			1250,
		},
	}

	for _, test := range table {
		result := test.input.GetAssetPrice()
		if result != test.expected {
			t.Errorf(
				"Transaction.GetAssetPrice(%f) = %f, expected %f",
				test.input.AssetPrice,
				result,
				test.expected,
			)
		}
	}
}