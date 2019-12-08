package downloader

import (
	"io"
)

//ProgressWritter interface
type ProgressWritter interface {
	Write(p []byte) (n int, err error)
	TotalBytes() int
	WrittenBytes() int
	ProgressPercent() chan float64
	CurrentProgress() float64
}

type progressWritter struct {
	totalBytes   int
	writtenBytes int
	writer       io.Writer
	progressChan chan float64
}

func newProgressWritter(totalBytes int, writer io.Writer) ProgressWritter {
	return &progressWritter{
		totalBytes:   totalBytes,
		writer:       writer,
		writtenBytes: 0,
		progressChan: make(chan float64, 500),
	}
}

//Write function
func (s *progressWritter) Write(p []byte) (n int, err error) {
	n, err = s.writer.Write(p)
	if err != nil {
		return
	}
	s.writtenBytes += len(p)
	return
}

//TotalBytes function
func (s *progressWritter) TotalBytes() int {
	return s.totalBytes
}

//WrittenBytes function
func (s *progressWritter) WrittenBytes() int {
	return s.writtenBytes
}

//ProgressPercent function
func (s *progressWritter) ProgressPercent() chan float64 {
	return s.progressChan
}

//CurrentProgress function
func (s *progressWritter) CurrentProgress() float64 {
	if s.TotalBytes() > 0 {
		return float64(s.WrittenBytes())
	}
	return float64(s.WrittenBytes()) / float64(s.TotalBytes())
}
