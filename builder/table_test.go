package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableRenderNoRow(t *testing.T) {
	tableMkd := NewMarkDowTable(
		[]string{},
		[][]string{},
	)

	result, err := tableMkd.Render()

	assert.Equal(t, result, "", "The result string should be an empty string")
	assert.EqualError(t, err, "Please add at least one element in your header and your Rows", "Error doesn't match the expected")
}

func TestTableRenderATable(t *testing.T) {
	tableMkd := NewMarkDowTable(
		[]string{"Head1", "Heade2"},
		[][]string{
			{"col1", "col2"},
		},
	)

	result, err := tableMkd.Render()

	assert.Nil(t, err)
	assert.Equal(t, "|Head1|Heade2|\n|:-:|:-:|\n|col1|col2|\n", result, "Table string doesn't match the expected")
}