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
		expected 	error
	} {
		{
			input: input{
				ctype: "link",
				content: "blabl",
			},
			expected: errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",")),
		},
		{
			input: input{
				ctype: "h9",
				content: "blabl",
			},
			expected: errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",")),
		},
		{
			input: input{
				ctype: "h1",
				content: "blabl",
			},
			expected: nil,
		},
	}

	for _, test := range table {
		result, err := NewMarkDowText(test.input.content, test.input.ctype)

		if 	(err != nil && test.expected == nil) ||
			(err == nil && test.expected != nil) ||
			(err != nil && test.expected != nil && err.Error() != test.expected.Error()) {

			t.Error(
				"Error trying to create a MarkDownText",
				"Input:",
				test.input,
				"Result",
				result,
			)

			if err != nil {
				t.Error(
					"Err",
					err.Error(),
				)
			}

			if test.expected != nil {
				t.Error(
					"Expected",
					test.expected.Error(),
				)
			}
		}

		if _, ok := result.(MarkDownBuilder); !ok && err == nil {
			t.Error(
				"Result doesn't have the same value than expected",
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

func TestRender(t *testing.T) {
	type input struct {
		ctype 	string
		content string
	}

	table := []struct {
		input 		input
		expected 	string
	} {
		{
			input: input{
				ctype: "h1",
				content: "blabl",
			},
			expected: "# blabl",
		},
		{
			input: input{
				ctype: "h5",
				content: "blabl",
			},
			expected: "##### blabl",
		},
		{
			input: input{
				ctype: "h2",
				content: "blabl",
			},
			expected: "## blabl",
		},
	}

	for _, test := range table {
		mkText, err := NewMarkDowText(test.input.content, test.input.ctype)

		if err != nil {
			t.Error(
				"Error Creating the MarkDowTextBuilder",
				err.Error(),
				"Input:",
				test.input,
			)
		}

		result, err := mkText.Render()
		if err != nil {
			t.Error(
				"Error Rendering the MarkDowTextBuilder",
				err.Error(),
				"Input:",
				test.input,
			)
		}

		if result != test.expected {
			t.Error(
				"Expected result from Render doesn't",
				err.Error(),
				"Input:",
				test.input,
			)
		}
	}
}