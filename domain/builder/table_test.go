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

	err := tableMkd.Render()

	assert.EqualError(t, err, "Please add at least one element in your header and your Rows", "Error doesn't match the expected")
}

func TestTableBuildATable(t *testing.T) {
	tableMkd := markDownTable{
		header: []string{"Head1", "Heade2"},
		rows: [][]string{
			{"col1", "col2"},
		},
	}

	result := tableMkd.Build()

	assert.Equal(t, "|Head1|Heade2|\n|:-:|:-:|\n|col1|col2|\n", result, "Table string doesn't match the expected")
}