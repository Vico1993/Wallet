package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutUnitPrice(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(100), operation.UnitPrice)
}

func TestWithProfitOutput(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, []string{
		"BTC",
		"0.1",
		"100",
		"10",
		"2632.5",
		"26225.07%",
	}, operation.WithProfit())
}

func TestWithProfitOutputWithEmptyCurrentPrice(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"COINSUPERCOIN",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, []string{
		"COINSUPERCOIN",
		"0.1",
		"100",
		"10",
		"0",
		"-100%",
	}, operation.WithProfit())
}

func TestGetCurrentPriceWith0AsCurrentAssetPrice(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(0), operation.GetCurrentPrice(0))
}

func TestGetCurrentPrice(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(10), operation.GetCurrentPrice(100))
}

func TestGetCurrentUnitPriceNoPriceFound(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"COINSUPERCOIN",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(0), operation.getCurrentUnitPrice())
}

func TestGetCurrentUnitPriceWithPriceFound(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(26325.07), operation.getCurrentUnitPrice())
}