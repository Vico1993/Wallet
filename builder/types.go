package builder

type MarkDownBuilder interface {
	Render() (string, error)
}