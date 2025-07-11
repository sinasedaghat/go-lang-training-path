package iomanager

type IOManager interface {
	InputReader() ([]string, error)
	OutputWriter(data any) error
}
