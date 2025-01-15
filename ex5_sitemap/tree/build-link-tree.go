package tree

type SitemapNode struct {
	url      string
	children []*SitemapNode
	parent   *SitemapNode
}

func BuildLinkTree(
	scrapePage ScrapePageFunc,
	startingUrl string,
) *SitemapNode {
	return traverse(
		scrapePage,
		startingUrl,
		nil,
		make(map[string]*SitemapNode),
	)
}

// TODO: TCO -- does it work here?
func traverse(
	scrapePage ScrapePageFunc,
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
			children = append(children, traverse(
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
