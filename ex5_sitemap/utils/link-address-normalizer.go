package utils

import (
	"net/url"
	pagescrapper "sitemap/page-scrapper"
)

func LinkAddressNormalizerFactory(baseUrl string) pagescrapper.NormalizeLinkAddressFunc {
	return func(urlAddress string) (string, error) {
		parsedUrl, err := url.Parse(urlAddress)
		if err != nil {
			return "", err
		}

		if parsedUrl.Host == "" {
			return url.JoinPath(baseUrl, parsedUrl.Path)
		}

		return url.JoinPath(urlAddress)
	}
}
