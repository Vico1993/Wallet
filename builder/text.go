package builder

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	TITLE = "title"
	ITALIC = "italic"
	ERROR = "error"
)

type markDownText struct {
	content string
	cType 	string
}

func (t markDownText) parseType() string {
	matchTitle, _ := regexp.MatchString("^h[1-6]$", t.cType)
	if (matchTitle) {
		return TITLE
	}

	if (strings.ToLower(t.cType) == ITALIC) {
		return ITALIC
	}

	return ""
}

func (t markDownText) validationOfType() error {
	tpe := t.parseType()

	if tpe != "" {
		return nil
	}

	return errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",") )
}

func NewMarkDowText(content string, ctype string) MarkDownBuilder {
	return &markDownText{
		cType: ctype,
		content: content,
	}
}

func (t markDownText) Render() (string, error) {
	err := t.validationOfType()

	if err != nil {
		return "", err
	}

	var renderString string

	switch t.parseType() {
    case TITLE:
		titleNumber, err := strconv.Atoi(
			strings.ReplaceAll(
				t.cType,
				"h",
				"",
			),
		)

		if err != nil {
			return "", err
		}

        renderString = strings.Repeat(
			"#",
			titleNumber,
		) + " " + t.content
    case ITALIC:
		renderString = "__" + t.content + "__"
	}

	return renderString + "\n", nil
}

func getSupportedType() []string {
	return []string{"h1", "h2", "h3", "h4", "h5", "h6"}
}