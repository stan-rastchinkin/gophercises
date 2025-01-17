package tree

import "sitemap/tree/queue"

func BuildLinkTreeDFS(
	scrapePage ScrapePageFunc,
	startingUrl string,
) *SitemapNode {
	return traverseDFS(
		scrapePage,
		startingUrl,
		nil,
		make(map[string]*SitemapNode),
	)
}

func traverseDFS(
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
			children = append(children, traverseDFS(
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

func BuildLinkTreeBFS(
	scrapePage ScrapePageFunc,
	startingUrl string,
) *SitemapNode {
	return traverseBFS(
		scrapePage,
		startingUrl,
		make(map[string]*SitemapNode),
	)
}

func traverseBFS(
	scrapePage ScrapePageFunc,
	url string,
	processedLinksRegistry map[string]*SitemapNode,
) *SitemapNode {
	nodeQueue := queue.New[*SitemapNode]()

	rootNode := SitemapNode{
		url:    url,
		parent: nil,
	}
	nodeQueue.Push(&rootNode)

	// TODO: sometimes doesn't make it to the last node somehow
	for {
		node := nodeQueue.Pull()
		if node == nil {
			break
		}

		if _, alreadyProcessed := processedLinksRegistry[node.url]; alreadyProcessed {
			break
		}

		for _, childLink := range scrapePage(node.url) {
			childNode := &SitemapNode{
				url:    childLink,
				parent: node,
			}

			node.children = append(node.children, childNode)
			nodeQueue.Push(childNode)
		}

		processedLinksRegistry[node.url] = node
	}

	return &rootNode
}
