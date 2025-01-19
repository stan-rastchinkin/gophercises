package tree

import "sitemap/tree/queue"

func BuildLinkTreeDFS(
	pageScrapper PageScrapper,
	startingUrl string,
) *SitemapNode {
	return traverseDFS(
		pageScrapper,
		startingUrl,
		make(map[string]*SitemapNode),
	)
}

// Uses built-in stack
func traverseDFS(
	pageScrapper PageScrapper,
	url string,
	processedLinks map[string]*SitemapNode,
) *SitemapNode {
	normalizedScrappedLinks := pageScrapper.GetLinks(url)

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
				pageScrapper,
				nestedlink,
				processedLinks,
			))
		}

	}
	currentNode.children = children

	return &currentNode
}

func BuildLinkTreeBFS(
	pageScrapper PageScrapper,
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

		for _, containedLink := range pageScrapper.GetLinks(node.url) {
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
