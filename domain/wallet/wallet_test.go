package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wallet = NewWallet([]Operation{
	NewOperation(
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
	),
	NewOperation(
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
	),
}, "test")

func TestGetTotalForUnitWithOnlyPurchase(t *testing.T) {
	result, err := wallet.GetQuantityByUnit("BTC")

	assert.Nil(t, err)
	assert.Equal(
		t,
		0.2,
		result,
	)
}

func TestOrderingByDate(t *testing.T) {
	var wallet = NewWallet([]Operation{
		NewOperation(
			"2022-06-25",
			0.1,
			"BTC",
			0,
			"fiat",
			10.0,
			10.0,
			"CAD",
			PURCHASE,
			"test",
		),
		NewOperation(
			"2022-06-15",
			0.1,
			"BTC",
			0,
			"fiat",
			10.0,
			10.0,
			"CAD",
			PURCHASE,
			"test",
		),
	}, "test")

	operations := wallet.GetOperations()

	assert.Equal(t, 2, len(operations))
	assert.Equal(t, "2022-06-15", operations[0].Date)
	assert.Equal(t, "2022-06-25", operations[1].Date)

	wallet.AddOperation(NewOperation(
		"2022-06-01",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	))
	operations = wallet.GetOperations()

	assert.Equal(t, 3, len(operations))
	assert.Equal(t, "2022-06-01", operations[0].Date)
	assert.Equal(t, "2022-06-25", operations[2].Date)
}

func TestGetTotalForUnitWithNoUnitRequested(t *testing.T) {
	result, err := wallet.GetQuantityByUnit("EGLD")

	assert.EqualError(t, err, "Unit not found in the wallet", "Error doesn't match the expected")
	assert.Equal(
		t,
		float64(0),
		result,
	)
}

func TestGetTotalForUnitWithUnitButExchangeLater(t *testing.T) {
	newWallet := NewWallet([]Operation{
		NewOperation(
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
		),
		NewOperation(
			"2022-06-19",
			1,
			"ETH",
			0,
			"BTC",
			0.05,
			100.0,
			"CAD",
			EXCHANGE,
			"test",
		),
	}, "test")

	resultETH, err := newWallet.GetQuantityByUnit("ETH")

	assert.Nil(t, err)
	assert.Equal( t, float64(1), resultETH )

	resultBTC, err := newWallet.GetQuantityByUnit("BTC")

	assert.Nil(t, err)
	assert.Equal( t, 0.05, resultBTC )
}

func TestGetTotalForUnitWithUnitButExchangeItLater(t *testing.T) {
	newWallet := NewWallet([]Operation{
		NewOperation(
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
		),
		NewOperation(
			"2022-06-19",
			1,
			"ETH",
			0,
			"BTC",
			0.1,
			100.0,
			"CAD",
			EXCHANGE,
			"test",
		),
	}, "test")

	resultETH, err := newWallet.GetQuantityByUnit("ETH")

	assert.Nil(t, err)
	assert.Equal( t, float64(1), resultETH )

	resultBTC, err := newWallet.GetQuantityByUnit("BTC")

	assert.Nil(t, err)
	assert.Equal( t, float64(0), resultBTC )
}

func TestGetTotalForUnitWithAfterAnAddOperation(t *testing.T) {
	result, err := wallet.GetQuantityByUnit("EGLD")

	assert.EqualError(t, err, "Unit not found in the wallet", "Error doesn't match the expected")
	assert.Equal(
		t,
		float64(0),
		result,
	)

	wallet.AddOperation(
		NewOperation(
			"2022-06-19",
			1,
			"EGLD",
			0,
			"fiat",
			10.0,
			10.0,
			"CAD",
			PURCHASE,
			"test",
		),
	)

	result, err = wallet.GetQuantityByUnit("EGLD")
	assert.Nil(t, err)
	assert.Equal(t, float64(1), result)
}