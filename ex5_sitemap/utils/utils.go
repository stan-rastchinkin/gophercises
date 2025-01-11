package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

// urlAddress must contain the whole url with host
func GetReaderFromWebPage(urlAddress string) (io.ReadCloser, error) {
	resp, err := http.Get(urlAddress)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// urlAddress must contain the whole url with host
func GetReaderFromLocalFs(urlAddress string) (io.ReadCloser, error) {
	fullPathToFile, err := urlAddressToFilePath(urlAddress)
	if err != nil {
		return nil, err
	}

	reader, err := os.OpenFile(fullPathToFile, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func urlAddressToFilePath(urlAddress string) (string, error) {
	u, err := url.Parse(urlAddress)
	if err != nil {
		return "", err
	}
	fmt.Printf("\n%+v\n", u.Path)

	basepath := path.Join("test-pages", u.Host)
	pathToFile := "index.html"
	if u.Path != "" && u.Path != "/" {
		pathToFile = u.Path + ".html"
	}

	return path.Join(basepath, pathToFile), nil
}
