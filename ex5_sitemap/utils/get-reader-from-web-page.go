package utils

import (
	"io"
	"net/http"
)

// urlAddress must contain the whole url with host
func GetReaderFromWebPage(urlAddress string) (io.ReadCloser, error) {
	resp, err := http.Get(urlAddress)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
