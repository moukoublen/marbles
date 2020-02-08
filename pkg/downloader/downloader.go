package downloader

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func filenameFromURL(url string) string {
	a := strings.Split(url, "/")
	if a == nil || len(a) == 0 {
		return ""
	}
	return a[len(a)-1]
}

var rgxFilename = regexp.MustCompile(`filename="(.*)"`)

func filenameFromContentDisposition(cd string) string {
	r := rgxFilename.FindStringSubmatch(cd)
	if r != nil && len(r) >= 2 {
		return r[1]
	}
	return ""
}

func filenameFromHeader(url string) string {
	resp, err := http.Head(url)
	if err != nil {
		return ""
	}

	n := resp.Header.Get("content-disposition")
	return n
}

// GetFilename returns the filename of file to download.
func GetFilename(url string) string {
	return ""
}

//Download function
func Download(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
