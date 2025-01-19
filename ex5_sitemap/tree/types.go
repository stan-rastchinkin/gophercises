package tree

type PageScrapper interface {
	GetLinks(string) []string
}

type SitemapNode struct {
	url      string
	children []*SitemapNode
	level    int
}
