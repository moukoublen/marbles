package downloader

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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

func newResponseWithContentDispositionHeader(filename string) *http.Response {
	return &http.Response{
		Header: newContentDispositionHeader(filename),
	}
}

func newResponseWithFileData(byteData []byte) *http.Response {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(byteData)),
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
			clientResponse{newResponseWithContentDispositionHeader("file1.tar.bz2"), nil},
			"file1.tar.bz2",
		},
		"content disposition contains filename 2": {
			"url3",
			clientResponse{newResponseWithContentDispositionHeader("file"), nil},
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

func TestDownload(t *testing.T) {
	var testcases = map[string]struct {
		url           string
		response      clientResponse
		expectedData  []byte
		expectedError error
	}{
		"client returns data": {
			"url1",
			clientResponse{newResponseWithFileData([]byte{0x01, 0x02}), nil},
			[]byte{0x01, 0x02},
			nil,
		},
		"client returns error": {
			"url1",
			clientResponse{&http.Response{}, errors.New("")},
			nil,
			errors.New(""),
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.url, func(tt *testing.T) {
			tt.Parallel()
			m := &MockHTTPClient{}
			m.Test(tt)
			m.On("Get", tc.url).Once().Return(tc.response.rsp, tc.response.err)
			d := NewDownloader(m)
			w := &bytes.Buffer{}
			err := d.Download(w, tc.url)
			if err != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				b := w.Bytes()
				assert.Equal(tt, tc.expectedData, b)
			}
		})
	}
}
