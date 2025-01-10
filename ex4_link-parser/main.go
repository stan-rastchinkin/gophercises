package main

import (
	"fmt"

	linkfilter "linkparser/filter-links"
	utils "linkparser/utils"
)

func main() {
	doc := utils.OpenAndParseHtmlFile("./filter-links/fixtures/ex4.html")

	links := linkfilter.FilterLinks(doc)

	for _, l := range links {
		fmt.Printf("l: %+v\n", *l)
	}
}
