package builder

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

type graph struct {
	data [][]float64
	options []asciigraph.Option
}

func NewGraph(opts []asciigraph.Option, data ...[]float64) MarkDownBuilder {
	return &graph{
		data: data,
		options: opts,
	}
}

func (g graph) Render() error {
	fmt.Println(
		asciigraph.PlotMany(
			g.data,
			g.options...,
		),
	)

	return nil
}
