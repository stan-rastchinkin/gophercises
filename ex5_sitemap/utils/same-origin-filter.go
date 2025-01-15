package utils

import (
	"net/url"
	pagescrapper "sitemap/page-scrapper"
)

func FilterSameOriginLinksFactory(baseUrl string) (pagescrapper.LinksFilterFunc, error) {
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
