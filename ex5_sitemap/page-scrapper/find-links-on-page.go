package pagescrapper

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	linkparser "linkparser/filter-links"
	tree "sitemap/tree"
)

// type PageScrapper struct {
// 	normalizer LinkNormalizer
// 	linkFilter LinkFilter
// 	readerAccessor
// }

func PageScrapperFactory(
	linkNormalizer LinkNormalizer,
	linkFilter LinkFilter,
	getReader GetReaderFunc,
) tree.ScrapePageFunc {

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
		linksOnPageSet := map[string]struct{}{}

		for _, link := range linkparser.FilterLinks(doc) {
			normalizedLink, err := linkNormalizer.Normalize(link.Href)
			handleErrorAndExit(err, fmt.Sprintf("Failed to normalize link %s", link.Href))
			isPassing, err := linkFilter.IsPassing(normalizedLink)
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
}

func handleErrorAndExit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		os.Exit(1)
	}
}
