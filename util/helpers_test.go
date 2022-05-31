package util

import (
	"testing"
)

func TestFormatFloat(t *testing.T) {
	table := []struct {
		input 		float64
		expected 	string
	}{
		{
			input: 3.467,
			expected: "3.46",
		},
		{
			input: 0.0067,
			expected: "0.0067",
		},
	}

	for _, test := range table {
		if FormatFloat(test.input) != test.expected {
			t.Errorf(
				"formatFloat(input: %f) - expected %s",
				test.input,
				test.expected,
			)
		}
	}
}