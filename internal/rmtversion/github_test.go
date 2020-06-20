package rmtversion

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type clientResponse struct {
	rsp *http.Response
	err error
}

func newResponse(data string) *http.Response {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(data))),
	}
}

// MockHTTPClient mock http client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	a := m.Called(url)
	return a.Get(0).(*http.Response), a.Error(1)
}

func TestGitHubGetLatestVersion(t *testing.T) {
	var testcases = map[string]struct {
		owner           string
		repo            string
		expectedURL     string
		response        clientResponse
		expectedVersion string
		expectError     bool
	}{
		"client returns error": {
			"owner",
			"repo",
			"https://api.github.com/repos/owner/repo/releases/latest",
			clientResponse{nil, errors.New("")},
			"",
			true,
		},
		"client retunrs json": {
			"owner",
			"repo",
			"https://api.github.com/repos/owner/repo/releases/latest",
			clientResponse{newResponse(`{"tag_name":"1.10.1"}`), nil},
			"1.10.1",
			false,
		},
	}

	for name, testcase := range testcases {
		tc := testcase
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			m := &MockHTTPClient{}
			m.Test(tt)
			m.On("Get", tc.expectedURL).Once().Return(tc.response.rsp, tc.response.err)
			g := RemoteVersionGitHub{
				OwnerName: tc.owner,
				RepoName:  tc.repo,
			}
			v, err := g.GetLatestVersion(m)
			if tc.expectError != (err != nil) {
				tt.Errorf("Expected error was %v and it is %v", tc.expectError, (err != nil))
			}
			if !tc.expectError && (v.String() != tc.expectedVersion) {
				tt.Errorf("Expected version is %s got %s", tc.expectedVersion, v.String())
			}
		})
	}
}
