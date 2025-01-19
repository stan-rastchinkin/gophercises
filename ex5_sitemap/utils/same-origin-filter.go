package utils

import (
	"net/url"
)

type SameOriginLinkFilter struct {
	ParsedBaseUrl *url.URL
}

func (filter SameOriginLinkFilter) IsPassing(urlAddress string) (bool, error) {
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		return false, err
	}

	return parsedUrl.Host == filter.ParsedBaseUrl.Host, nil
}

func NewSameOriginLinkFilter(baseUrl string) (SameOriginLinkFilter, error) {
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return SameOriginLinkFilter{}, err
	}

	return SameOriginLinkFilter{ParsedBaseUrl: parsedBaseUrl}, nil
}
