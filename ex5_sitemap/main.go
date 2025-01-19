package main

import (
	"fmt"
	"os"

	ps "sitemap/page-scrapper"
	tree "sitemap/tree"
	"sitemap/utils"
)

type ProgramConfig struct {
	getReader   ps.GetReaderFunc
	siteHomeUrl string
}

func main() {
	sitemap := program(&ProgramConfig{
		getReader:   utils.GetReaderFromLocalFs,
		siteHomeUrl: "https://www.test.org/",
	})

	fmt.Println("\nResult:")
	fmt.Print(tree.RenderToXml(sitemap))
}

func program(config *ProgramConfig) *tree.SitemapNode {
	linkNormalizer := utils.LinkNormalizer{BaseUrl: config.siteHomeUrl}
	sameOriginFilter, err := utils.NewSameOriginLinkFilter(config.siteHomeUrl)
	handleErrorAndExit(err, "Failed to create same origin filter")

	scrapeLinksOnPage := ps.PageScrapperFactory(
		linkNormalizer,
		sameOriginFilter,
		config.getReader,
	)

	return tree.BuildLinkTreeBFS(
		scrapeLinksOnPage,
		config.siteHomeUrl,
	)
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
