package utils

import "io"

type GetReaderFunc func(urlAddress string) (io.ReadCloser, error)
