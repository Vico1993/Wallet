package builder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextRenderTitle1(t *testing.T) {
	mkText := NewMarkDowText("blabl", "h1")
	result, err := mkText.Render()

	assert.Nil(t, err)
	assert.Equal(t, "# blabl\n", result, "The markdown string is incorrect")
}

func TestTextRenderTitle5(t *testing.T) {
	mkText := NewMarkDowText("blabl", "h5")
	result, err := mkText.Render()

	assert.Nil(t, err)
	assert.Equal(t, result, "##### blabl\n", "The markdown string is incorrect")
}

func TestTextRenderTitle2(t *testing.T) {
	mkText := NewMarkDowText("blabl", "h2")
	result, err := mkText.Render()

	assert.Nil(t, err)
	assert.Equal(t, result, "## blabl\n", "The markdown string is incorrect")
}

func TestTextRenderItalic(t *testing.T) {
	mkText := NewMarkDowText("blabl", "italic")
	result, err := mkText.Render()

	assert.Nil(t, err)
	assert.Equal(t, result, "__blabl__\n", "The markdown string is incorrect")
}

func TestTextRenderItalicCapitalType(t *testing.T) {
	mkText := NewMarkDowText("blabl", "iTAlic")
	result, err := mkText.Render()

	assert.Nil(t, err)
	assert.Equal(t, result, "__blabl__\n", "The markdown string is incorrect")
}

func TestTextRenderTypeNotSupported(t *testing.T) {
	mkText := NewMarkDowText("blabl", "link")
	result, err := mkText.Render()

	assert.Equal(t, result, "", "The result string should be an empty string")
	assert.EqualError(t, err, "Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ","), "Error doesn't match the expected")
}