package utils

import "net/url"

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
