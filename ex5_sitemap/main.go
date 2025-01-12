package main

import (
	"fmt"
	"io"
	"os"

	linkparser "linkparser/filter-links"

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
		fmt.Printf("link: %v\n", link)
	}
}

func program(config *ProgramConfig) []string {
	normalizeLink := utils.LinkAddressNormalizerFactory(config.siteHomeUrl)
	sameOriginFilter, err := utils.FilterSameOriginLinksFactory(config.siteHomeUrl)
	handleErrorAndExit(err, "Failed to create same origin filter")

	return findLinksOnPage(
		normalizeLink,
		sameOriginFilter,
		config.siteHomeUrl,
		config.getReader,
	)
}

func findLinksOnPage(
	normalizeLinkAddress utils.NormalizeLinkAddressFunc,
	sameOriginFilter utils.FilterSameOriginLinksFunc,
	// TODO: I'm passing the whole config here - bad abstractions
	pageUrl string,
	getReader GetReaderFunc,
) []string {

	pageReader, err := getReader(pageUrl)
	// TODO: modify local test pages to contain a finite graph (no links to absent pages)
	handleErrorAndExit(err, "Failed to fetch page")
	defer pageReader.Close()

	doc, err := html.Parse(pageReader)
	handleErrorAndExit(err, "Failed to parse page")

	// todo: cover with tests duplicate case
	// remove duplicates
	sameOriginLinksOnPageSet := map[string]struct{}{}

	for _, link := range linkparser.FilterLinks(doc) {
		normalizedLink, err := normalizeLinkAddress(link.Href)
		handleErrorAndExit(err, fmt.Sprintf("Failed to normalize link %s", link.Href))
		isSameOrigin, err := sameOriginFilter(normalizedLink)
		handleErrorAndExit(err, fmt.Sprintf("Same origin filter failed on link %s", normalizedLink))
		if isSameOrigin {
			sameOriginLinksOnPageSet[normalizedLink] = struct{}{}
		}
	}

	var result []string
	for link := range sameOriginLinksOnPageSet {
		result = append(result, link)
	}

	return result
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		if len(msg) != 0 {
			fmt.Printf("%s: %e", msg, err)
		} else {
			fmt.Printf("%e", err)
		}
		os.Exit(1)
	}
}
