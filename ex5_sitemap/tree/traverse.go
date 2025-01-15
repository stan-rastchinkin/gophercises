package tree

import (
	pagescrapper "sitemap/page-scrapper"
)

type SitemapNode struct {
	url      string
	children []*SitemapNode
	parent   *SitemapNode
}

// TODO: TCO -- does it work here?
func Traverse(
	scrapePage pagescrapper.ScrapePageFunc,
	url string,
	parentNode *SitemapNode,
	processedLinksRegistry map[string]*SitemapNode,
) *SitemapNode {
	normalizedScrappedLinks := scrapePage(url)

	currentNode := SitemapNode{
		url:    url,
		parent: parentNode,
	}

	processedLinksRegistry[url] = &currentNode

	var children []*SitemapNode
	for _, nestedlink := range normalizedScrappedLinks {
		if existingNode, alreadyProcessed := processedLinksRegistry[nestedlink]; alreadyProcessed {
			children = append(children, existingNode)
		} else {
			children = append(children, Traverse(
				scrapePage,
				nestedlink,
				&currentNode,
				processedLinksRegistry,
			))
		}

	}
	currentNode.children = children

	return &currentNode
}
