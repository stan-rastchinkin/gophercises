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
		siteHomeUrl: "http://www.iana.org",
	})

	for _, link := range links {
		fmt.Printf("links: %v\n", *link)
	}
}

func program(config *ProgramConfig) []*filterlinks.Link {
	pageReader, err := config.getReader(config.siteHomeUrl)
	if err != nil {
		fmt.Printf("Failed to fetch page: %e", err)
		os.Exit(1)
	}
	defer pageReader.Close()

	doc, err := html.Parse(pageReader)
	if err != nil {
		fmt.Printf("Failed to parse response body: %e", err)
		os.Exit(1)
	}

	return filterlinks.FilterLinks(doc)
}
