package downloader

import (
	"testing"
)

type mockWriter struct {
	WriteFunc func(p []byte) (n int, err error)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return m.WriteFunc(p)
}

func TestProgressWriterSuccess(t *testing.T) {
	m := &mockWriter{
		WriteFunc: func(p []byte) (n int, err error) {
			return len(p), nil
		},
	}

	s := newProgressWritter(0, m)

	chunkOne := []byte{1, 2, 3, 4, 5, 6}
	expectedLen := len(chunkOne)

	n, _ := s.Write(chunkOne)
	if s.WrittenBytes() != expectedLen {
		t.Errorf("Error in shadow writer. Written bytes expected to be %d got %d", expectedLen, s.WrittenBytes())
	}
	if n != expectedLen {
		t.Errorf("Error in shadow writer. Written bytes returned from Write function expected to be %d got %d", expectedLen, n)
	}
}
