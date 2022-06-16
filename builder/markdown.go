package builder

// TODO: Improve this file
type MarkDownBuilder interface {
	Render() 	error
}

type MarkDownData struct {
	Block 		MarkDownBuilder
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

		err = element.Block.Render()
		if err != nil {
			return err
		}
	}

	return nil
}