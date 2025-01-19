package tree

import (
	"encoding/xml"
	"fmt"
	"os"
	"sitemap/tree/queue"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

type loc struct {
	Value string `xml:"loc"`
}

func RenderToXml(entryNode *SitemapNode) string {
	processedUrls := map[string]struct{}{}
	nodeQueue := queue.New[*SitemapNode]()

	toXml := urlset{
		Xmlns: xmlns,
	}

	nodeQueue.Push(entryNode)

	for {
		node, isEmpty := nodeQueue.Pull()
		if isEmpty {
			break
		}
		if _, isProcessed := processedUrls[node.url]; isProcessed {
			continue
		}

		toXml.Urls = append(toXml.Urls, loc{node.url})
		for _, child := range node.children {
			nodeQueue.Push(child)
		}

		processedUrls[node.url] = struct{}{}
	}

	rendered, err := xml.MarshalIndent(toXml, "", "  ")
	if err != nil {
		fmt.Printf("Failed to render tree into XML: %e\n", err)
		os.Exit(1)
	}

	return xml.Header + string(rendered)
}
