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

func TestIsInStringSlice(t *testing.T) {
	type input struct {
		str 	string
		list 	[]string
	}

	table := []struct {
		input 		input
		expected 	bool
	}{
		{
			input: input{
				str: "toto",
				list: []string{"test", "word", "not", "in"},
			},
			expected: false,
		},
		{
			input: input{
				str: "in",
				list: []string{"test", "word", "not", "in"},
			},
			expected: true,
		},
	}

	for _, test := range table {
		result := IsInStringSlice(test.input.str, test.input.list) != test.expected
		if result {
			t.Error(
				"IsInStringSlice doesn't return the expected result",
				"input:",
				test.input,
				"expected:",
				test.expected,
				"result:",
				result,
			)
		}
	}
}
