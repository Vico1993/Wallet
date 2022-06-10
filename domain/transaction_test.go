package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAssetPriceOneTransaction(t *testing.T) {
	input := Transaction{
		Price: 100,
		Date: "2022-05-26",
		Quantity: 1,
		Asset: "BTC",
		AssetPrice: 100,
	}

	assert.Equal(t, float64(100), input.GetAssetPrice(), "Asset price should be 100")
}

func TestGetAssetPriceWithNoAssetPriceInTransaction(t *testing.T) {
	input := Transaction{
		Price: 100,
		Date: "2022-05-26",
		Quantity: 1,
		Asset: "BTC",
	}

	assert.Equal(t, float64(100), input.GetAssetPrice(), "Asset price should be 100")
}

func TestGetAssetPriceWithNoAssetPriceInTransactionAndCalculQuatity(t *testing.T) {
	input := Transaction{
		Price: 100,
		Date: "2022-05-26",
		Quantity: 0.8,
		Asset: "BTC",
	}

	assert.Equal(t, float64(125), input.GetAssetPrice(), "Asset price should be 100")
}

func TestGetAssetPriceWithNoAssetPriceInTransactionAndSmallQuatity(t *testing.T) {
	input := Transaction{
		Price: 100,
		Date: "2022-05-26",
		Quantity: 0.08,
		Asset: "BTC",
	}

	assert.Equal(t, float64(1250), input.GetAssetPrice(), "Asset price should be 100")
}