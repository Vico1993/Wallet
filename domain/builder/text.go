package builder

import (
	"Vico1993/Wallet/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	TITLE = "title"
	ITALIC = "italic"
	TEXT = "text"
	LIST = "list"
	ERROR = "error"
)

type markDownText struct {
	content string
	cType 	string
	data 	[]interface{}
}

func (t markDownText) parseType() string {
	matchTitle, _ := regexp.MatchString("^h[1-6]$", t.cType)
	if (matchTitle) {
		return TITLE
	}

	if (strings.ToLower(t.cType) == ITALIC) {
		return ITALIC
	}

	if (t.cType == "text") {
		return TEXT
	}

	if (t.cType == "list") {
		return LIST
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

func NewMarkDowText(content string, ctype string, data []interface{}) MarkDownBuilder {
	return &markDownText{
		cType: ctype,
		content: content,
		data: data,
	}
}

func (t markDownText) Build() (string, error) {
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
	case TEXT:
		renderString = t.content
	case LIST:
		comaMatch := regexp.MustCompile(",")
		if comaMatch.FindString(t.content) != "" {
			for _, item := range strings.Split(t.content, ",") {
				renderString += "- " + item + " \n"
			}
		} else {
			renderString = "- " + t.content
		}
	}

	paramMatch := regexp.MustCompile("%s")
	if paramMatch.FindString(renderString) != "" && t.data != nil {
		renderString = fmt.Sprintf(
			renderString,
			t.data...,
		)
	}

	return renderString, nil
}

func (t markDownText) Render() error {
	err := t.validationOfType()

	if err != nil {
		return err
	}

	renderString, err := t.Build()

	if err!= nil {
		return err
	}

	return util.RenderMarkdown(renderString)
}

func getSupportedType() []string {
	return []string{"h1", "h2", "h3", "h4", "h5", "h6", "text", "list"}
}