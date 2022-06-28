package builder

// TODO: Improve this file
type MarkDownBuilder interface {
	Render() 	error
}

type Data struct {
	Block 		MarkDownBuilder
}

type MarkDown struct {
	data []Data
}

func NewMarkDown(d []Data) MarkDownBuilder {
	return &MarkDown{
		data: d,
	}
}

func (m *MarkDown) AddData(d Data) {
	m.data = append(m.data, d)
}

func (m MarkDown) Render() error {
	for _, element := range m.data {
		err := element.Block.Render()
		if err != nil {
			return err
		}
	}

	return nil
}