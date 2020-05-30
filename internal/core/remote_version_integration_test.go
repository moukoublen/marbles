// +build integration

package core

import (
	"net/http"
	"testing"
)

func TestGetLatestRelease(t *testing.T) {
	var testcases = map[string]struct {
		owner string
		repo  string
	}{
		"docker compose": {"docker", "compose"},
		"docker machine": {"docker", "machine"},
	}

	for name, testcase := range testcases {
		tc := testcase
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			g := RemoteVersionGithub{
				OwnerName: tc.owner,
				RepoName:  tc.repo,
			}
			v, err := g.GetLatestVersion(http.DefaultClient)
			if err != nil {
				t.Errorf("Error during version fetch %s", err.Error())
			}
			tt.Logf("Latest version for %s/%s: %#v", tc.owner, tc.repo, v.String())
		})
	}
}
