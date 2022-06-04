package builder

import (
	"errors"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	type input struct {
		ctype 	string
		content string
	}

	table := []struct {
		input 		input
		expected 	string
		err 		error
	} {
		{
			input: input{
				ctype: "h1",
				content: "blabl",
			},
			expected: "# blabl\n",
			err: nil,
		},
		{
			input: input{
				ctype: "h5",
				content: "blabl",
			},
			expected: "##### blabl\n",
			err: nil,
		},
		{
			input: input{
				ctype: "h2",
				content: "blabl",
			},
			expected: "## blabl\n",
			err: nil,
		},
		{
			input: input{
				ctype: "link",
				content: "blabl",
			},
			expected: "",
			err: errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",") ),
		},
	}

	for _, test := range table {
		mkText := NewMarkDowText(test.input.content, test.input.ctype)

		result, err := mkText.Render()
		if err != nil && err.Error() != test.err.Error() {
			t.Error(
				"Unexpected error of Render with error: ",
				err.Error(),
				"Expected error:",
				test.err,
				"Input:",
				test.input,
			)
		}

		if result != test.expected {
			t.Error(
				"Expected result from Render doesn't match:",
				test.expected,
				"Result:",
				result,
				"Input:",
				test.input,
			)
		}
	}
}