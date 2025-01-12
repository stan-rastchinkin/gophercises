package tree

import "fmt"

func RenderToString(node *SitemapNode, level uint) string {
	result := fmt.Sprintf("level %d:\n", level) + node.url + "\n\n"

	for _, child := range node.children {
		result += RenderToString(child, level+1)
	}

	return result
}
