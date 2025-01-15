package utils

import (
	"net/url"
	pagescrapper "sitemap/page-scrapper"
)

func LinkAddressNormalizerFactory(baseUrl string) pagescrapper.NormalizeLinkAddressFunc {
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
