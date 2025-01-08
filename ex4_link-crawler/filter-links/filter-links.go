package filterlinks

import (
	"slices"

	html "golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func FilterLinks(parsedTree *html.Node) []*Link {
	var links []*Link

	for n := range parsedTree.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			hrefAttribute := n.Attr[slices.IndexFunc(
				n.Attr,
				func(el html.Attribute) bool {
					return el.Key == "href"
				})]

			links = append(links, &Link{
				Href: hrefAttribute.Val,
				Text: n.FirstChild.Data,
			})
		}
	}

	return links
}
