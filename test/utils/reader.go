package utils

import (
	"io"
)

// Reader struct.
type Reader struct {
	chunks [][]byte
	index  int
}

// NewReader constructor
func NewReader(chunks [][]byte) *Reader {
	return &Reader{
		chunks: chunks,
		index:  0,
	}
}

// Size returns the final size of all data.
func (r *Reader) Size() int {
	s := 0
	for _, chunk := range r.chunks {
		s += len(chunk)
	}
	return s
}

// NumOfChunks returns the number of chunks.
func (r *Reader) NumOfChunks() int {
	return len(r.chunks)
}

// Read function.
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.index == len(r.chunks) {
		err = io.EOF
		return
	}
	n = copy(p, r.chunks[r.index])
	r.index++
	if r.index == len(r.chunks) {
		err = io.EOF
	}
	return
}
