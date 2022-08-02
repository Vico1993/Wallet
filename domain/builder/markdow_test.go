package builder

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdownAddData(t *testing.T) {
	var markdown MarkDown

	markdown.AddData(
		Data{
			Block: NewMarkDowText("We are missing the csv file path", "h1", nil),
		},
	)

	assert.Equal(t, 1, len(markdown.data))

	markdown.AddData(
		Data{
			Block: NewMarkDowText("We are missing the csv file path", "h1", nil),
		},
	)

	assert.Equal(t, 2, len(markdown.data))
}

func TestMarkdownRenderWithError(t *testing.T) {
	var markdown MarkDown

	markdown.AddData(
		Data{
			Block: NewMarkDowText("We are missing the csv file path", "newSuperList", nil),
		},
	)

	err := markdown.Render()
	assert.EqualError(
		t,
		errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",") ),
		err.Error(),
	)
}

func TestMarkdownRenderWithoutError(t *testing.T) {
	var markdown MarkDown

	markdown.AddData(
		Data{
			Block: NewMarkDowText("We are missing the csv file path", "text", nil),
		},
	)

	err := markdown.Render()
	assert.Nil(t, err)
}