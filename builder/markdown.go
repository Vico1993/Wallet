package builder

// TODO: Improve this class
type MarkDownBuilder interface {
	Render() 	error
}

type MarkDownData struct {
	String 		MarkDownBuilder
}

type MarkDown struct {
	data []MarkDownData
}

func NewMarkDown(d []MarkDownData) MarkDownBuilder {
	return &MarkDown{
		data: d,
	}
}

func (m MarkDown) Render() error {
	for _, element := range m.data {
		var err error

		err = element.String.Render()
		if err != nil {
			return err
		}
	}

	return nil
}