package rmtversion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/moukoublen/marbles/internal/core"
)

// RemoteVersionGitHub remote version functionality for github
type RemoteVersionGitHub struct {
	OwnerName string
	RepoName  string
}

const githubLatestReleaseF = "https://api.github.com/repos/%s/%s/releases/latest"

// GetLatestVersion from github
func (g RemoteVersionGitHub) GetLatestVersion(client HTTPClient) (*core.Version, error) {
	url := fmt.Sprintf(githubLatestReleaseF, g.OwnerName, g.RepoName)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil {
		return nil, err
	}

	v, err := core.ParseVersion(keyAsString(m, "tag_name"))
	if err == nil {
		return v, nil
	}

	v, err = core.ParseVersion(keyAsString(m, "name"))
	return v, err
}

func keyAsString(j map[string]interface{}, key string) string {
	if i, found := j[key]; found {
		if o, isString := i.(string); isString {
			return o
		}
	}
	return ""
}
