package builder

import (
	"errors"
	"strings"
	"testing"
)

func TestValidationFailed(t *testing.T) {
	type input struct {
		ctype 	string
		content string
	}

	table := []struct {
		input 		input
		expected 	*markDownText
		err 		error
	} {
		{
			input: input{
				ctype: "link",
				content: "blabl",
			},
			expected: nil,
			err: errors.New("Type not supported at the moment, only support: " + strings.Join(title, " | ")),
		},
		{
			input: input{
				ctype: "h1",
				content: "blabl",
			},
			expected: &markDownText{
				cType: "h1",
				content: "blabl",
			},
			err: nil,
		},
	}

	for _, test := range table {
		result, err := NewMarkDowText(test.input.content, test.input.ctype)

		if err != test.err && result != test.expected {
			t.Error(
				"Error trying to create a MarkDownText",
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