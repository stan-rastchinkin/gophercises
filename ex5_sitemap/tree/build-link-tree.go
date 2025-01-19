package tree

import "sitemap/tree/queue"

func BuildLinkTreeDFS(
	scrapePage ScrapePageFunc,
	startingUrl string,
) *SitemapNode {
	return traverseDFS(
		scrapePage,
		startingUrl,
		make(map[string]*SitemapNode),
	)
}

// Uses built-in stack
func traverseDFS(
	scrapePage ScrapePageFunc,
	url string,
	processedLinks map[string]*SitemapNode,
) *SitemapNode {
	normalizedScrappedLinks := scrapePage(url)

	currentNode := SitemapNode{
		url: url,
	}

	processedLinks[url] = &currentNode

	var children []*SitemapNode
	for _, nestedlink := range normalizedScrappedLinks {
		if existingNode, alreadyProcessed := processedLinks[nestedlink]; alreadyProcessed {
			children = append(children, existingNode)
		} else {
			children = append(children, traverseDFS(
				scrapePage,
				nestedlink,
				processedLinks,
			))
		}

	}
	currentNode.children = children

	return &currentNode
}

func BuildLinkTreeBFS(
	scrapePage ScrapePageFunc,
	startUrl string,
) *SitemapNode {
	nodeQueue := queue.New[*SitemapNode]()
	processedUrls := map[string]struct{}{}

	rootNode := &SitemapNode{
		url:      startUrl,
		children: []*SitemapNode{},
	}
	nodeQueue.Push(rootNode)

	for {
		node, isEmpty := nodeQueue.Pull()
		if isEmpty {
			break
		}

		if _, exists := processedUrls[node.url]; exists {
			continue
		}

		for _, containedLink := range scrapePage(node.url) {
			childNode := &SitemapNode{
				url:      containedLink,
				children: []*SitemapNode{},
			}
			node.children = append(node.children, childNode)

			nodeQueue.Push(childNode)
		}

		processedUrls[node.url] = struct{}{}
	}

	return rootNode
}
