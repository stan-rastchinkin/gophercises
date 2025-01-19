package pagescrapper

import "io"

type LinkNormalizer interface {
	Normalize(string) (string, error)
}

type LinkFilter interface {
	IsPassing(string) (bool, error)
}

type GetReaderFunc func(urlAddress string) (io.ReadCloser, error)
