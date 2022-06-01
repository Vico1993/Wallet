package builder

import (
	"errors"
	"testing"
)

func TestRenderTable(t *testing.T) {
	type input struct {
		header 	[]string
		rows 	[][]string
	}

	table := []struct {
		input 		input
		expected 	string
		err 		error
	} {
		{
			input: input{
				header: []string{},
				rows: [][]string{},
			},
			expected: "",
			err: errors.New("Please add at least one element in your header and your Rows"),
		},
		{
			input: input{
				header: []string{"Head1", "Heade2"},
				rows: [][]string{
					{"col1", "col2"},
				},
			},
			expected: "| Head1 | Heade2 |\n| :-: |:-: |\n|col1|col2|\n",
			err: nil,
		},
	}

	for _, test := range table {
		tableMkd := NewMarkDowTable(
			test.input.header,
			test.input.rows,
		)

		result, err := tableMkd.Render()

		if err != test.err && result != test.expected {
			t.Error(
				"Error trying to render a Table",
				"Input:",
				test.input,
				"Result",
				result,
				"Expected:",
				test.expected,
			)
		}
	}
}