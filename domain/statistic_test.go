package domain

import (
	"reflect"
	"testing"
)

func TestCalculProfit(t *testing.T) {
	type input struct {
		start 	float64
		end 	float64
	}

	table := []struct {
		input 		input
		expected 	float64
	}{
		{
			input{
				start: 50,
				end: 50,
			},
			0,
		},
		{
			input{
				start: 50,
				end: 100,
			},
			100,
		},
		{
			input{
				start: 50,
				end: 10,
			},
			-80,
		},
		{
			input{
				start: 1,
				end: 10,
			},
			900,
		},
	}

	for _, test := range table {
		result := calculProfit(test.input.start, test.input.end)
		if result != test.expected {
			t.Errorf(
				"CalculProfit(Start: %f, End: %f) = %f, expected %f",
				test.input.start,
				test.input.end,
				result,
				test.expected,
			)
		}
	}
}

func TestAddInvest(t *testing.T) {
	type input struct {
		symbol 		string
		invest 		float64
		value 		float64
		quantity 	float64
	}

	table := []struct {
		input 		[]input
		expected 	Statistic
	}{
		{
			input: append(
				make([]input, 1),
				input{
					symbol: "BTC",
					invest: 10,
					value: 100,
				},
			),
			expected: Statistic{
				invest: 10,
				value: 100,
			},
		},
		{
			input: append(
				make([]input, 2),
				input{
					symbol: "BTC",
					invest: 10,
					value: 100,
					quantity: 1,
				},
				input{
					symbol: "BTC",
					invest: 100,
					value: 1000,
					quantity: 2,
				},
			),
			expected: Statistic{
				invest: 110,
				value: 1100,
			},
		},
		{
			input: append(
				make([]input, 3),
				input{
					symbol: "BTC",
					invest: 10,
					value: 100,
					quantity: 1,
				},
				input{
					symbol: "BTC",
					invest: 100,
					value: 1000,
					quantity: 2,
				},
				input{
					symbol: "ETH",
					invest: 50,
					value: 500,
					quantity: 0.1,
				},
			),
			expected: Statistic{
				invest: 160,
				value: 1600,
			},
		},
	}

	for _, test := range table {
		var result = Statistic{}

		for _, i := range test.input {
			result.AddInvest(i.symbol, i.invest, i.value, i.quantity)
		}

		if result.invest != test.expected.invest && result.value != test.expected.value && len(result.details) != len(test.input) {
			t.Errorf(
				"AddInvest input(Invest: %f, Value: %f, len(details): %d) = expected(Invest: %f, Value: %f, len(details): %d)",
				result.invest,
				result.value,
				len(result.details),
				test.expected.invest,
				test.expected.value,
				len(test.expected.details),
			)
		}
	}
}

func TestGetDetails(t *testing.T) {
	type input struct {
		symbol 		string
		invest 		float64
		value 		float64
		quantity 	float64
	}

	table := []struct {
		input 		[]input
		expected 	[]Details
	} {
		{
			input: 	[]input{
				{
					invest: 100,
					value: 50,
					symbol: "BTC",
					quantity: 1,
				},
				{
					invest: 100,
					value: 50,
					symbol: "BTC",
					quantity: 1,
				},
				{
					invest: 10,
					value: 100,
					symbol: "ETH",
					quantity: 1,
				},
			},
			expected: []Details {
				{
					Symbol: "ETH",
					Profit: 900,
					Quantity: 1,
				},
				{
					Symbol: "BTC",
					Profit: -50,
					Quantity: 2,
				},
			},
		},
	}

	for _, test := range table {
		var stats = Statistic{}

		for _, i := range test.input {
			stats.AddInvest(i.symbol, i.invest, i.value, i.quantity)
		}

		if !reflect.DeepEqual(stats.GetDetails(), test.expected) {
			t.Error(
				"getDetails return incorrect value than expected",
				"\nGetDetails Value:\n",
				stats.GetDetails(),
				"\nInput:\n",
				test.input,
				"\nExpected:\n",
				test.expected,
			)
		}
	}
}