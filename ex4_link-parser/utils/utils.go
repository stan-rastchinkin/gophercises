package utils

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func OpenAndParseHtmlFile(path string) *html.Node {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		os.Exit(1)
	}

	return doc
}
