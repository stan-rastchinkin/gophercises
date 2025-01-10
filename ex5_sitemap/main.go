package main

import (
	"fmt"
	"net/http"
	"os"

	filterlinks "linkparser/filter-links"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://www.iana.org/")
	if err != nil {
		fmt.Printf("Failed to fetch page: %e", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Failed to parse response body: %e", err)
		os.Exit(1)
	}

	links := filterlinks.FilterLinks(doc)

	for _, link := range links {
		fmt.Printf("links: %v\n", *link)
	}
}
