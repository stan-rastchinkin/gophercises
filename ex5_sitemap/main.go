package main

import (
	"fmt"
	"io"
	"os"

	ps "sitemap/page-scrapper"
	tree "sitemap/tree"
	"sitemap/utils"
)

type ProgramConfig struct {
	siteHomeUrl string
	getReader   func(string) (io.ReadCloser, error)
	maxDepth    int8
}

func main() {
	sitemapTree := program(&ProgramConfig{
		siteHomeUrl: "https://www.test.org/",
		getReader:   utils.GetReaderFromLocalFs,
		maxDepth:    10,
	})

	// fmt.Println("\nXML Result:")
	// fmt.Print(tree.RenderToXml(sitemapTree))
	fmt.Println("\nResult:")
	fmt.Print(tree.RenderToString(sitemapTree))
}

func program(config *ProgramConfig) *tree.SitemapNode {
	sameOriginFilter, err := utils.NewSameOriginLinkFilter(config.siteHomeUrl)
	handleErrorAndExit(err, "Failed to create same origin filter")

	scrapper := ps.PageScrapper{
		LinkNormalizer: utils.LinkNormalizer{BaseUrl: config.siteHomeUrl},
		LinkFilter:     sameOriginFilter,
		GetPageReader:  config.getReader,
	}

	return tree.BuildLinkTreeBFS(
		scrapper,
		config.siteHomeUrl,
		config.maxDepth,
	)
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
