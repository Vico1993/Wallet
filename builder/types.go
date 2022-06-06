package builder

type MarkDownBuilder interface {
	Render() (string, error)
}

type MasterBuilder struct {
	String 	interface{}
	Data 	[]interface{}
}