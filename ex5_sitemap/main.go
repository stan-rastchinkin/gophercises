package main

import (
	"fmt"
	"io"
	"os"

	filterlinks "linkparser/filter-links"

	"sitemap/utils"

	"golang.org/x/net/html"
)

type GetReaderFunc func(urlAddress string) (io.ReadCloser, error)

type ProgramConfig struct {
	getReader   GetReaderFunc
	siteHomeUrl string
}

func main() {
	links := program(&ProgramConfig{
		getReader:   utils.GetReaderFromLocalFs,
		siteHomeUrl: "https://www.iana.org/",
	})

	for _, link := range links {
		fmt.Printf("links: %v\n", *link)
	}
}

func program(config *ProgramConfig) []*filterlinks.Link {
	normalizeLinkAddress := utils.LinkAddressNormalizerFactory(config.siteHomeUrl)
	sameOriginFilter, err := utils.FilterSameOriginLinksFactory(config.siteHomeUrl)
	handleError(err, "Failed to create same origin filter")

	pageReader, err := config.getReader(config.siteHomeUrl)
	handleError(err, "Failed to fetch page")
	defer pageReader.Close()

	doc, err := html.Parse(pageReader)
	handleError(err, "Failed to parse response body")

	return filterlinks.FilterLinks(doc)
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
