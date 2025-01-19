package pagescrapper

type LinkNormalizer interface {
	Normalize(string) (string, error)
}

type LinkFilter interface {
	IsPassing(string) (bool, error)
}
