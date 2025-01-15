package pagescrapper

import "io"

type LinksFilterFunc func(urlAddress string) (bool, error)
type NormalizeLinkAddressFunc func(urlAddress string) (string, error)
type GetReaderFunc func(urlAddress string) (io.ReadCloser, error)
