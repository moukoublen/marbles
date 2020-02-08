package utils

// Writer struct
type Writer struct {
	WriteFunc func(p []byte) (int, error)
}

func (m Writer) Write(p []byte) (n int, err error) {
	return m.WriteFunc(p)
}

// DropWriter just returns successful writes
var DropWriter = Writer{
	WriteFunc: func(p []byte) (int, error) {
		return len(p), nil
	},
}
