// Package downloader continas download functionality bundled with file name resolving.
package downloader

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

// ContentDispositionHeader header name
const ContentDispositionHeader = "content-disposition"

// HTTPClient is the interface for http client methods
type HTTPClient interface {
	Head(url string) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
}

// Downloader structure
type Downloader struct {
	Client HTTPClient
}

// NewDownloader Downloader constructor
func NewDownloader(client HTTPClient) *Downloader {
	return &Downloader{
		Client: client,
	}
}

// Download the file from the given url in the given file path.
func (d *Downloader) Download(writer io.Writer, url string) error {
	resp, err := d.Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(writer, resp.Body)
	return err
}

// GetFilename returns the filename of file to download.
func (d *Downloader) GetFilename(url string) string {
	f := d.filenameFromHeader(url)
	if len(f) > 0 {
		return f
	}
	return filenameFromURL(url)
}

var rgxFilename = regexp.MustCompile(`filename="(.*)"`)

func (d *Downloader) filenameFromHeader(url string) string {
	resp, err := d.Client.Head(url)
	if err != nil {
		return ""
	}

	cd := resp.Header.Get(ContentDispositionHeader)
	r := rgxFilename.FindStringSubmatch(cd)
	if len(r) >= 2 {
		return r[1]
	}

	return ""
}

func filenameFromURL(url string) string {
	a := strings.Split(url, "/")
	if len(a) == 0 {
		return ""
	}
	return a[len(a)-1]
}
