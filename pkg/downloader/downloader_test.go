package downloader

import (
	"testing"
)

func TestGetFilenameFromURL(t *testing.T) {
	var testcases = []struct {
		url              string
		expectedFilename string
	}{
		{"", ""},
		{"https://download.internet.o/files/file1.tar.bz2", "file1.tar.bz2"},
		{"https://download.internet.o/files/", ""},
		{"file", "file"},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.url, func(t *testing.T) {
			t.Parallel()
			n := filenameFromURL(tc.url)
			if n != tc.expectedFilename {
				t.Errorf("getFilenameFromURL failed for url %s. Expected %s got %s", tc.url, tc.expectedFilename, n)
			}
		})
	}
}

func TestFilenameFromContentDisposition(t *testing.T) {
	var testcases = []struct {
		url              string
		expectedFilename string
	}{
		{`filename="filename.jpg"`, "filename.jpg"},
		{`nope`, ""},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.url, func(t *testing.T) {
			t.Parallel()
			n := filenameFromContentDisposition(tc.url)
			if n != tc.expectedFilename {
				t.Errorf("getFilenameFromURL failed for url %s. Expected %s got %s", tc.url, tc.expectedFilename, n)
			}
		})
	}
}
