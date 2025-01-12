package pagescrapper

import (
	"fmt"
	"os"
	"sitemap/utils"

	"golang.org/x/net/html"

	linkparser "linkparser/filter-links"
)

type ScrapePageFunc func(pageUrl string) []string

func PageScrapperFactory(
	normalizeLinkAddress utils.NormalizeLinkAddressFunc,
	sameOriginFilter utils.FilterSameOriginLinksFunc,
	getReader utils.GetReaderFunc,
) ScrapePageFunc {

	return func(
		pageUrl string,
	) []string {
		fmt.Printf("Srapping link: %s\n", pageUrl)

		pageReader, err := getReader(pageUrl)
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
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
