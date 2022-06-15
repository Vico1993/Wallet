package builder

import (
	"errors"
	"fmt"
	"regexp"
)

type MarkDownBuilder interface {
	Render() (string, error)
}

type MarkDownData struct {
	String 		interface{}
	Variable 	[]interface{}
}

type MarkDown struct {
	data []MarkDownData
}

func NewMarkDown(d []MarkDownData) MarkDownBuilder {
	return &MarkDown{
		data: d,
	}
}

func (m MarkDown) Render() (string, error) {
	render := ""

	for _, element := range m.data {
		var err error
		var renderStr string

		if s, ok := element.String.(string); ok {
			renderStr = s
		} else if s, ok := element.String.(MarkDownBuilder); ok {
			renderStr, err = s.Render()
			if err != nil {
				return "", err
			}
		} else {
			return "", errors.New("Type not supported")
		}

		paramMatch := regexp.MustCompile("%s")
		if paramMatch.FindString("%s") != "" {
			render += fmt.Sprintf(
				renderStr,
				element.Variable...,
			)
		} else {
			render += fmt.Sprint(
				renderStr,
			)
		}
	}

	return render, nil
}