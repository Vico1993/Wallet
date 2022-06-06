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

// @todo: Add test
func IsInStringSlice(a string, list []string) bool {
	for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// @todo: Add test
func TransformStringSliceIntoInterface(list []string) []interface{} {
	vals := make([]interface{}, len(list))
    for i, v := range list {
		vals[i] = v
	}

	return vals
}