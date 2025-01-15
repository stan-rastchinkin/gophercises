package tree

import (
	"fmt"
	"strconv"
)

type RegistryLog struct {
	node  *SitemapNode
	level uint
}

func RenderToString(node *SitemapNode) string {
	return renderToString(
		node,
		0,
		make(map[string]RegistryLog),
	)
}

func renderToString(
	node *SitemapNode,
	level uint,
	processedNodesRegestry map[string]RegistryLog,
) string {
	if processedNode, nodeAlreadyProcessed := processedNodesRegestry[node.url]; nodeAlreadyProcessed {
		return fmt.Sprintf("level %d:\n", level) +
			node.url +
			"\nNote: This Node is already rendered at level " +
			strconv.FormatUint(uint64(processedNode.level), 10)
	}

	result := fmt.Sprintf("level %d:\n", level) + node.url + "\n\n"

	for _, child := range node.children {
		result += renderToString(child, level+1, processedNodesRegestry)
	}

	processedNodesRegestry[node.url] = RegistryLog{
		node:  node,
		level: level,
	}

	return result
}
