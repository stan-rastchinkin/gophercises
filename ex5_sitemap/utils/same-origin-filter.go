package utils

import (
	"net/url"
)

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
