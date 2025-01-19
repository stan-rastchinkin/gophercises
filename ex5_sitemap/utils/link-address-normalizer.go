package utils

import (
	"net/url"
)

type LinkNormalizer struct {
	BaseUrl string
}

func (normalizer LinkNormalizer) Normalize(urlAddress string) (string, error) {
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		return "", err
	}

	if parsedUrl.Host == "" {
		return url.JoinPath(normalizer.BaseUrl, parsedUrl.Path)
	}

	return url.JoinPath(urlAddress)
}
