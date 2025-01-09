package filterlinks

import (
	"regexp"
	"slices"
	"strings"

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
				Text: getNodeInnerText(n),
			})
		}
	}

	return links
}

func getNodeInnerText(node *html.Node) string {
	var fragments []string

	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			fragments = append(fragments, n.Data)
		}
	}

	concatenated := strings.Join(fragments, "")
	re := regexp.MustCompile(`\W*(?P<result>.*)\W*`)
	matches := re.FindStringSubmatch(concatenated)
	groups := re.SubexpNames()

	idx := slices.IndexFunc(groups, func(el string) bool { return el == "result" })

	return matches[idx]
}
