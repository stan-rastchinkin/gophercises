package utils

import (
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

	basepath := path.Join("test-pages", u.Host)
	pathToFile := "index.html"
	if u.Path != "" && u.Path != "/" {
		pathToFile = u.Path + ".html"
	}

	return path.Join(basepath, pathToFile), nil
}

type NormalizeLinkAddressFunc func(urlAddress string) (string, error)

func LinkAddressNormalizerFactory(baseUrl string) NormalizeLinkAddressFunc {
	return func(urlAddress string) (string, error) {
		u, err := url.Parse(urlAddress)
		if err != nil {
			return "", err
		}

		if u.Host == "" {
			return url.JoinPath(baseUrl, u.Path)
		}

		return urlAddress, nil
	}
}

type FilterSameOriginLinksFunc func(urlAddress string) (bool, error)

func FilterSameOriginLinksFactory(baseUrl string) (FilterSameOriginLinksFunc, error) {
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	productFunc := func(urlAddress string) (bool, error) {
		parsedUrl, err := url.Parse(urlAddress)
		if err != nil {
			return false, err
		}

		return parsedUrl.Host == parsedBaseUrl.Host, nil
	}

	return productFunc, nil
}
