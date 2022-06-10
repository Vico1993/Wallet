package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculProfitWithNoGain(t *testing.T) {
	result := calculProfit(50, 50)

	assert.Equal(t, float64(0), result, "Result should be equal to 0")
}

func TestCalculProfitGain(t *testing.T) {
	result := calculProfit(50, 100)

	assert.Equal(t, float64(100), result, "Result should be equal to 100")
}

func TestCalculProfitWithNegativeGain(t *testing.T) {
	result := calculProfit(50, 10)

	assert.Equal(t, float64(-80), result, "Result should be equal to -80")
}

func TestCalculProfitWithLargeGain(t *testing.T) {
	result := calculProfit(1, 10)

	assert.Equal(t, float64(900), result, "Result should be equal to 900")
}

func TestAddInvestOneTransaction(t *testing.T) {
	var result = Statistic{}

	result.AddInvest("BTC", 10, 100, 1)

	assert.Equal(
		t,
		Statistic{
			invest: 10,
			value: 100,
			details: map[string]dStatistic{
				"BTC": {
					invest: 10,
					value: 100,
					quantity: 1,
				},
			},
		},
		result,
		"Struct doesn't match the expected",
	)
}

func TestAddInvestTwoTransactionsSameCoin(t *testing.T) {
	var result = Statistic{}

	result.AddInvest("BTC", 10, 100, 1)
	result.AddInvest("BTC", 100, 1000, 2)

	assert.Equal(
		t,
		Statistic{
			invest: 110,
			value: 1100,
			details: map[string]dStatistic{
				"BTC": {
					invest: 110,
					value: 1100,
					quantity: 3,
				},
			},
		},
		result,
		"Struct doesn't match the expected",
	)
}

func TestAddInvestTwoCoin(t *testing.T) {
	var result = Statistic{}

	result.AddInvest("BTC", 10, 100, 1)
	result.AddInvest("BTC", 100, 1000, 2)
	result.AddInvest("ETH", 50, 500, 0.1)

	assert.Equal(
		t,
		Statistic{
			invest: 160,
			value: 1600,
			details: map[string]dStatistic{
				"BTC": {
					invest: 110,
					value: 1100,
					quantity: 3,
				},
				"ETH": {
					invest: 50,
					value: 500,
					quantity: 0.1,
				},
			},
		},
		result,
		"Struct doesn't match the expected",
	)
}

func TestDetailsWithTwoCoin(t *testing.T) {
	var result = Statistic{}

	result.AddInvest("BTC", 10, 50, 1)
	result.AddInvest("BTC", 100, 50, 1)
	result.AddInvest("ETH", 10, 100, 1)

	assert.Equal(
		t,
		[]Details {
			{
				Symbol: "ETH",
				Profit: 900,
				Quantity: 1,
			},
			{
				Symbol: "BTC",
				Profit: -9.09,
				Quantity: 2,
			},
		},
		result.GetDetails(),
		"Details doesn't match the expectation",
	)
}
