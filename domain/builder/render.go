package builder

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

// TODO: Find out why we have a BIG spaces
func renderMarkdown(s string) error {
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width
		glamour.WithWordWrap(100),
	)

	out, err := r.Render(
		strings.ReplaceAll(s, "\t", ""),
	)

	if err != nil {
		return err
	}

	// Don't print anything in TEST MODE
	if (os.Getenv("TEST") != "1") {
		fmt.Print(out)
	}

	return nil
}