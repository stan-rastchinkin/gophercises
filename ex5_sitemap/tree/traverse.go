package tree

import (
	pagescrapper "sitemap/page-scrapper"
)

type SitemapNode struct {
	url      string
	children []*SitemapNode
	parent   *SitemapNode
}

func Traverse(
	pageScrapper pagescrapper.ScrapePageFunc,
	url string,
	parentNode *SitemapNode,
) *SitemapNode {
	normalizedScrappedLinks := pageScrapper(url)

	currentNode := SitemapNode{
		url:    url,
		parent: parentNode,
	}

	var children []*SitemapNode
	for _, nestedlink := range normalizedScrappedLinks {
		children = append(children, Traverse(
			pageScrapper,
			nestedlink,
			&currentNode,
		))
	}
	currentNode.children = children

	return &currentNode
}
