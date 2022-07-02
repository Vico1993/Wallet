package builder

import (
	"Vico1993/Wallet/util"
	"errors"
	"strings"
)

type markDownTable struct {
	header 			[]string
	rows 			[][]string
	renderString 	string
}

func NewMarkDowTable(header []string, rows [][]string) MarkDownBuilder {
	return &markDownTable{
		header: header,
		rows: rows,
	}
}

func (t *markDownTable) buildHeader() {
	t.renderString += "|" + strings.Join(t.header, "|") + "|" + "\n"
	t.renderString += "|" + strings.Repeat(":-:|", len(t.header)) + "\n"
}

func (t *markDownTable) buildRows() {
	for _, row := range t.rows {
		t.renderString += "|" + strings.Join(row, "|") + "|\n"
	}
}

func (t markDownTable) Build() (string) {
	// First the Header
	t.buildHeader()

	// Then Rows
	t.buildRows()

	return t.renderString
}

func (t markDownTable) Render() error {
	if len(t.header) == 0 || len(t.rows) == 0 {
		return errors.New("Please add at least one element in your header and your Rows")
	}

	return util.RenderMarkdown(t.Build())
}
