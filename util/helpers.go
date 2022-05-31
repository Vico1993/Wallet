package util

import (
	"math"
	"strconv"
)

func FormatFloat(numb float64) string {
	round := math.Floor(numb * 100) / 100

	if round != 0 {
		numb = round
	}

	return strconv.FormatFloat(numb, 'g', -1, 64)
}