package main

import (
	"fmt"

	linkfilter "linkcrawler/filter-links"
	utils "linkcrawler/utils"
)

func main() {
	doc := utils.OpenAndParseHtmlFile("./test/ex1.html")

	links := linkfilter.FilterLinks(doc)

	for _, l := range links {
		fmt.Printf("l: %v\n", *l)
	}
}
