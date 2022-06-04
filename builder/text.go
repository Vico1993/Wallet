package builder

import (
	"errors"
	"regexp"
	"strings"
)

type markDownText struct {
	content string
	cType 	string
}

func (t markDownText) validationOfType() (string, error) {
	matchTitle, _ := regexp.MatchString("^h[1-6]$", t.cType)
	if (matchTitle) {
		return "title", nil
	}

	return "error", errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",") )
}

func NewMarkDowText(content string, ctype string) (markDownText, error) {
	t := markDownText{
		cType: ctype,
		content: content,
	}

	_, err := t.validationOfType()

	if err != nil {
		return markDownText{}, err
	}

	return t, nil
}


func getSupportedType() []string {
	return []string{"h1", "h2", "h3", "h4", "h5", "h6"}
}