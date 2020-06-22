package io

import (
	"io"
	"math"
	"testing"
	"time"

	"github.com/moukoublen/marbles/test/utils"
)

func TestProgressSingleWriterSuccess(t *testing.T) {
	chunkOne := []byte{0x1, 0x2, 0x3}
	expectedLen := len(chunkOne)

	m := utils.DropWriter
	s := NewProgressWritterWithInternalChannel("test", expectedLen, m)

	n, _ := s.Write(chunkOne)

	if s.writtenBytes != expectedLen {
		t.Errorf("Error in shadow writer. Written bytes expected to be %d got %d", expectedLen, s.writtenBytes)
	}

	if n != expectedLen {
		t.Errorf("Error in shadow writer. Written bytes returned from Write function expected to be %d got %d", expectedLen, n)
	}

	r := <-s.Progress()

	if !r.Complete() {
		t.Errorf("Error in progress writer. Should be complete")
	}
}

func TestProgressWriterWithInternalChannel(t *testing.T) {
	r := utils.NewReader([][]byte{
		{0x1, 0x2, 0x3},
		{0x1, 0x2, 0x3, 0x4},
		{0x1, 0x2, 0x3, 0x4, 0x5},
		{0x1, 0x2, 0x3, 0x4},
		{0x1, 0x2, 0x3},
	})

	s := NewProgressWritterWithInternalChannel("test", r.Size(), utils.DropWriter)

	_, _ = io.Copy(s, r)

	p := receiveProgressReport(t, s.Progress())
	assertCompletion(t, 0.157, p.Completion)
	p = receiveProgressReport(t, s.Progress())
	assertCompletion(t, 0.368, p.Completion)
	p = receiveProgressReport(t, s.Progress())
	assertCompletion(t, 0.631, p.Completion)
	p = receiveProgressReport(t, s.Progress())
	assertCompletion(t, 0.842, p.Completion)
	p = receiveProgressReport(t, s.Progress())
	assertCompletion(t, 1.000, p.Completion)

	if !p.Complete() {
		t.Fatalf("Progress should be completed and it is not")
	}

}

func receiveProgressReport(t *testing.T, ch <-chan WritterProgressUpdate) WritterProgressUpdate {
	select {
	case <-time.Tick(2 * time.Second):
		t.Error("Report did not arrive")
		return WritterProgressUpdate{}
	case report := <-ch:
		return report
	}
}

func assertCompletion(t *testing.T, a, b float64) {
	if math.Abs(a-b) > 1e-3 {
		t.Errorf("Completion float error. Expected %f is %f", a, b)
	}
}
