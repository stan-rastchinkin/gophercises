package pagescrapper

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"

	linkparser "linkparser/filter-links"
)

type PageScrapper struct {
	LinkNormalizer LinkNormalizer
	LinkFilter     LinkFilter
	GetPageReader  func(urlAddress string) (io.ReadCloser, error)
}

func (ps PageScrapper) GetLinks(pageUrl string) []string {
	fmt.Printf("Srapping link: %s\n", pageUrl)

	pageReader, err := ps.GetPageReader(pageUrl)
	handleErrorAndExit(err, "Failed to fetch page")
	defer pageReader.Close()

	doc, err := html.Parse(pageReader)
	handleErrorAndExit(err, "Failed to parse page")

	// todo: cover with tests duplicate case
	// remove duplicates
	linksOnPageSet := map[string]struct{}{}

	for _, link := range linkparser.FilterLinks(doc) {
		normalizedLink, err := ps.LinkNormalizer.Normalize(link.Href)
		handleErrorAndExit(err, fmt.Sprintf("Failed to normalize link %s", link.Href))
		isPassing, err := ps.LinkFilter.IsPassing(normalizedLink)
		handleErrorAndExit(err, fmt.Sprintf("Same origin filter failed on link %s", normalizedLink))
		if isPassing {
			linksOnPageSet[normalizedLink] = struct{}{}
		}
	}

	var result []string
	for link := range linksOnPageSet {
		result = append(result, link)
	}

	return result
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
