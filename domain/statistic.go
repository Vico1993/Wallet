package domain

import "math"

type Statistic struct {
	TotalInvest float64
	ActualValue float64
}

func (s Statistic) GetTotalProfit() (float64) {
	profit := ((s.ActualValue - s.TotalInvest) / s.TotalInvest) * 100

	return math.Round(profit*100)/100
}