package wallet

import "math"

func calculProfit(startValue float64, endValue float64) (float64) {
	profit := ((endValue - startValue) / startValue) * 100

	return math.Round(profit*100)/100
}
