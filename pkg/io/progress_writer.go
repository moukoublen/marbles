package io

import (
	"io"
	"math"
	"time"
)

// WritterProgressUpdate struct
type WritterProgressUpdate struct {
	Name       string
	Completion float64
	Timestamp  int64
}

// Complete check if completion percent is 100%.
func (w WritterProgressUpdate) Complete() bool {
	threshold := 1e-9
	return math.Abs(float64(1)-w.Completion) <= threshold
}

// ProgressWritter struct
type ProgressWritter struct {
	name         string
	totalBytes   int
	writtenBytes int
	writer       io.Writer
	pc           progressChannel
}

// NewProgressWritterWithInternalChannel creates a new ProgressWritter with internal created progress channel
func NewProgressWritterWithInternalChannel(name string, totalBytes int, writer io.Writer) *ProgressWritter {
	return &ProgressWritter{
		name:         name,
		totalBytes:   totalBytes,
		writtenBytes: 0,
		writer:       writer,
		pc: internalProgressChannel{
			ch: make(chan WritterProgressUpdate, 300),
		},
	}
}

// NewProgressWritterWithExternalChannel creates a new ProgressWritter given a progress channel
func NewProgressWritterWithExternalChannel(name string, totalBytes int, writer io.Writer, progressCh chan WritterProgressUpdate) *ProgressWritter {
	return &ProgressWritter{
		name:         name,
		totalBytes:   totalBytes,
		writtenBytes: 0,
		writer:       writer,
		pc: externalProgressChannel{
			ch: progressCh,
		},
	}
}

// Write function
func (s *ProgressWritter) Write(p []byte) (int, error) {
	n, err := s.writer.Write(p)
	if err != nil {
		return n, err
	}
	s.writtenBytes += len(p)
	s.reportProgress()
	if cpc, ok := s.pc.(io.Closer); ok && s.finish() {
		cpc.Close()
	}
	return n, nil
}

// ReadFrom uses io.Copy hidding the current method to avoid infinite loop.
func (s *ProgressWritter) ReadFrom(r io.Reader) (int64, error) {
	return io.Copy(struct{ io.Writer }{s}, r)
}

// CurrentProgressPersent function
func (s *ProgressWritter) CurrentProgressPersent() float64 {
	if s.totalBytes <= 0 {
		return float64(s.writtenBytes)
	}
	return float64(s.writtenBytes) / float64(s.totalBytes)
}

// Progress returns the receive channel that progress reports arrive to.
func (s *ProgressWritter) Progress() <-chan WritterProgressUpdate {
	return s.pc.progressChannel()
}

func (s *ProgressWritter) reportProgress() {
	w := WritterProgressUpdate{
		Name:       s.name,
		Completion: s.CurrentProgressPersent(),
		Timestamp:  time.Now().UTC().Unix(),
	}
	s.pc.progressChannel() <- w
}

func (s *ProgressWritter) finish() bool {
	return s.writtenBytes == s.totalBytes
}

//
// Progress Channel
//
type progressChannel interface {
	progressChannel() chan WritterProgressUpdate
}

type internalProgressChannel struct {
	ch chan WritterProgressUpdate
}

func (i internalProgressChannel) Close() error {
	close(i.ch)
	return nil
}

func (i internalProgressChannel) progressChannel() chan WritterProgressUpdate {
	return i.ch
}

type externalProgressChannel struct {
	ch chan WritterProgressUpdate
}

func (i externalProgressChannel) progressChannel() chan WritterProgressUpdate {
	return i.ch
}
