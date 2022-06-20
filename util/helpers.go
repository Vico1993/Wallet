package util

import (
	"log"
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

func IsInStringSlice(a string, list []string) bool {
	for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func TransformStringSliceIntoInterface(list []string) []interface{} {
	vals := make([]interface{}, len(list))
	for i, v := range list {
		vals[i] = v
	}

	return vals
}

func TransformStringToFloat(str string) float64 {
	flt, err := strconv.ParseFloat(str, 10)
	if (err != nil ) {
		log.Fatalln( "Error parsing Float: ", err.Error())
	}

	return flt
}

func ReverseSlice[S ~[]E, E any](s S) S {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

	return s
}