package util

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/guptarohit/asciigraph"
)

// TODO: Find out why we have a BIG spaces
func RenderMarkdown(s string) error {
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
	)

	out, err := r.Render(
		strings.ReplaceAll(s, "\t", ""),
	)
	if err != nil {
		return err
	}

	fmt.Print(out)

	return nil
}

func RenderGraph(lists ...[]float64) {
	fmt.Println(
		asciigraph.PlotMany(
			lists,
		),
	)
}