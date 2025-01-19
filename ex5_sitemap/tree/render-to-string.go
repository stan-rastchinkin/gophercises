package tree

import (
	"fmt"
	"sitemap/tree/queue"
)

func RenderToString(
	entryNode *SitemapNode,
) string {
	nodeQueue := queue.New[*SitemapNode]()
	processedUrls := map[string]struct{}{}

	nodeQueue.Push(entryNode)

	result := ""

	for {
		node, isEmpty := nodeQueue.Pull()
		if isEmpty {
			break
		}

		if _, exists := processedUrls[node.url]; exists {
			continue
		}

		result += fmt.Sprintf("url: %s\nchildren: ", node.url)
		for _, child := range node.children {
			result += child.url + "  "
			nodeQueue.Push(child)
		}
		result += "\n\n"

		processedUrls[node.url] = struct{}{}
	}

	return result
}
