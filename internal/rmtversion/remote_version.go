package rmtversion

import (
	"net/http"

	"github.com/moukoublen/marbles/internal/core"
)

// HTTPClient interface
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// RemoteVersion interface
type RemoteVersion interface {
	GetLatestVersion(client HTTPClient) (*core.Version, error)
}
