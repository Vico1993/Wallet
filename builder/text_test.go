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
			err: errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",")),
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

		if 	(err != nil && test.err == nil) ||
			(err == nil && test.err != nil) ||
			(err != nil && test.err != nil && err.Error() != test.err.Error()) {

			t.Error(
				"Error trying to create a MarkDownText",
				"Input:",
				test.input,
				"Result",
				result,
				"Expected:",
				test.expected,
			)

			if err!= nil {
				t.Error(
					"Err",
					err.Error(),
				)
			}

			if test.err != nil {
				t.Error(
					"Err",
					test.err.Error(),
				)
			}
		}

		if result != nil &&
			test.expected != nil &&
			result.cType != test.expected.cType &&
			result.content != test.expected.content {
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