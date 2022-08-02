package builder

import (
	"fmt"
	"os"

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
	// Don't print anything in TEST MODE
	if (os.Getenv("TEST") != "1") {
		fmt.Println(
			asciigraph.PlotMany(
				g.data,
				g.options...,
			),
		)
	}

	return nil
}
