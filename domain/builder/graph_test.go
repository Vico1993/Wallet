package builder

import (
	"testing"

	"github.com/guptarohit/asciigraph"
	"github.com/stretchr/testify/assert"
)

func TestHappyPathGraphRender(t *testing.T) {
	graph := NewGraph(
		[]asciigraph.Option{
			asciigraph.SeriesColors(
				asciigraph.Red,
				asciigraph.White,
			),
		},
		[]float64{4.3, 6.7, 12.9, 10.2},
		make([]float64, 4),
	)

	err := graph.Render()
	assert.Nil(t, err)
}