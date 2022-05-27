package domain

import (
	"fmt"
	"math"
	"sort"
)

type profit struct {
	invest float64
	value float64
}

type Statistic struct {
	invest float64
	value float64
	details map[string]profit
}

func calculProfit (startValue float64, endValue float64) (float64) {
	profit := ((endValue - startValue) / startValue) * 100

	return math.Round(profit*100)/100
}

func (s *Statistic) AddInvest(symbol string, invest float64, value float64) {
	s.value += value
	s.invest += invest

	// Initialisation of the map
	if (s.details == nil) {
		s.details = make(map[string]profit)
	}

	if entry, ok := s.details[symbol]; ok {
		entry.invest += invest
		entry.value += invest

		s.details[symbol] = entry
	} else {
		s.details[symbol] = profit{
			invest: invest,
			value: value,
		}
	}
}

func (s Statistic) GetTotalProfit() (float64) {
	return calculProfit(s.invest, s.value)
}

func (s Statistic) GetTotalInvest() (float64) {
	fmt.Println(s.invest)

	return s.invest
}

type Details struct {
	Symbol  string
	Profit  float64
}

func (s Statistic) GetDetails() ([]Details) {
	var detailsSorted []Details
	for symbol, stat := range s.details {
		profit := calculProfit(stat.invest, stat.value)

		detailsSorted = append(detailsSorted, Details{symbol, profit})
	}

	sort.Slice(detailsSorted, func(i, j int) bool {
        return detailsSorted[i].Profit > detailsSorted[j].Profit
    })

	return detailsSorted
}