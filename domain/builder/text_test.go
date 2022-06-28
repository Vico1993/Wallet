package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextBuildTitle1(t *testing.T) {
	mkText := markDownText{
		cType: "h1",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, result, "# blabl", "The markdown string is incorrect")
}

func TestTextBuildTitle5(t *testing.T) {
	mkText := markDownText{
		cType: "h5",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "##### blabl", result, "The markdown string is incorrect")
}

func TestTextBuildTitle2(t *testing.T) {
	mkText := markDownText{
		cType: "h2",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "## blabl", result, "The markdown string is incorrect")
}

func TestTextBuildItalic(t *testing.T) {
	mkText := markDownText{
		cType: "italic",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "__blabl__", result, "The markdown string is incorrect")
}

func TestTextBuildItalicCapitalType(t *testing.T) {
	mkText := markDownText{
		cType: "iTAlic",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "__blabl__", result, "The markdown string is incorrect")
}

func TestTextBuildText(t *testing.T) {
	mkText := markDownText{
		cType: "text",
		data: nil,
		content: "blabl",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "blabl", result, "The result string should be an empty string")
}

func TestTextBuildTextWithVariable(t *testing.T) {
	mkText := markDownText{
		cType: "text",
		data: []interface{}{"toto"},
		content: "blabl %s",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "blabl toto", result, "The result string should be an empty string")
}

func TestTextBuildTextWithVariableNueric(t *testing.T) {
	mkText := markDownText{
		cType: "text",
		data: []interface{}{7},
		content: "blabl %d",
	}
	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "blabl 7", result, "The result string should be an empty string")
}

func TestTextBuildList(t *testing.T) {
	mkText := markDownText{
		cType: "list",
		data: nil,
		content: "blabl,toto",
	}

	result, err := mkText.Build()

	assert.Nil(t, err)
	assert.Equal(t, "- blabl \n- toto \n", result, "The result string should be an empty string")
}