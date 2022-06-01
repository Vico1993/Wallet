package builder

import (
	"Vico1993/Wallet/util"
	"errors"
	"strings"
)

var title = []string{"h1", "h2", "h3"}

type markDownText struct {
	content string
	cType 	string
}

func (t markDownText) validationOfType() error {
	if !util.IsInStringSlice(t.cType, title) {
		return errors.New("Type not supported at the moment, only support: " + strings.Join(title, " | "))
	}

	return nil
}

func NewMarkDowText(content string, ctype string) (*markDownText, error) {
	t := markDownText{
		cType: ctype,
		content: content,
	}

	err := t.validationOfType()

	if err != nil {
		return nil, err
	}

	return &t, nil
}