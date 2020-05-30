package downloader

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type clientResponse struct {
	rsp *http.Response
	err error
}

func newContentDispositionHeader(filename string) http.Header {
	h := http.Header{}
	h.Set(ContentDispositionHeader, fmt.Sprintf(`filename="%s"`, filename))
	return h
}

func newResponse(filename string) *http.Response {
	return &http.Response{
		Header: newContentDispositionHeader(filename),
	}
}

// MockHTTPClient mock http client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Head(url string) (resp *http.Response, err error) {
	a := m.Called(url)
	return a.Get(0).(*http.Response), a.Error(1)
}

func (m *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	a := m.Called(url)
	return a.Get(0).(*http.Response), a.Error(1)
}

func TestGetFilename(t *testing.T) {
	var testcases = map[string]struct {
		url              string
		response         clientResponse
		expectedFilename string
	}{
		"client returns error": {
			"",
			clientResponse{nil, errors.New("")},
			"",
		},
		"content disposition contains filename 1": {
			"url2",
			clientResponse{newResponse("file1.tar.bz2"), nil},
			"file1.tar.bz2",
		},
		"content disposition contains filename 2": {
			"url3",
			clientResponse{newResponse("file"), nil},
			"file",
		},
		"name from url 1": {
			"https://download.internet.o/files/file1.tar.bz2",
			clientResponse{&http.Response{}, nil},
			"file1.tar.bz2",
		},
		"No name available": {
			"https://download.internet.o/files/",
			clientResponse{&http.Response{}, nil},
			"",
		},
		"edge case but why not": {
			"file",
			clientResponse{&http.Response{}, nil},
			"file",
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.url, func(tt *testing.T) {
			tt.Parallel()
			m := &MockHTTPClient{}
			m.Test(tt)
			m.On("Head", tc.url).Once().Return(tc.response.rsp, tc.response.err)
			d := NewDownloader(m)
			filename := d.GetFilename(tc.url)
			if filename != tc.expectedFilename {
				tt.Errorf("getFilenameFromURL failed for url. Expected \"%s\" got \"%s\"", tc.expectedFilename, filename)
			}
		})
	}

}
