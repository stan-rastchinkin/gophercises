package tree

type ScrapePageFunc func(pageUrl string) []string

type SitemapNode struct {
	url      string
	children []*SitemapNode
	parent   *SitemapNode
}
