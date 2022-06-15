package builder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataTypeNotSupported(t *testing.T) {
	markdown := NewMarkDown([]MarkDownData{
		{
			String: 56.2,
		},
	})

	str, err := markdown.Render()


	assert.Equal(t, errors.New("Type not supported"), err)
	assert.Equal(t, "", str)
}

func TestRenderMarkdownSuccess(t *testing.T) {
	markdown := NewMarkDown([]MarkDownData{
		{
			String: NewMarkDowText("Wallet", "h1"),
		},
	})

	str, err := markdown.Render()

	assert.Nil(t, err)
	assert.Equal(t, "# Wallet\n", str)
}