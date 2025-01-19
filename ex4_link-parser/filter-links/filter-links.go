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

			attrIndex := slices.IndexFunc(
				n.Attr,
				func(el html.Attribute) bool {
					return el.Key == "href"
				})
			if attrIndex == -1 {
				continue // corrupted link w/o href
			}
			hrefAttribute := n.Attr[attrIndex]

			links = append(links, &Link{
				Href: hrefAttribute.Val,
				Text: getNodeInnerText(n),
			})
		}
	}

	return links
}

func getNodeInnerText(node *html.Node) string {
	var concatenated string

	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			concatenated += n.Data
		}
	}

	return trimStringWithNativeTools(concatenated)
}

// Complecated and doesn't trim spaces at end of the group
func trimStringWithRegexGroup(input string) string {
	re := regexp.MustCompile(`\W*(?P<result>.+)\W*`)
	matches := re.FindStringSubmatch(input)
	groups := re.SubexpNames()

	idx := slices.IndexFunc(groups, func(el string) bool { return el == "result" })

	return matches[idx]
}

// Better regex solution - alternatives
// also \s == [\r\n\t\f\v]
func trimStringWithRegexAlternatives(input string) string {
	re := regexp.MustCompile(`^\s+|\s+$`)

	return re.ReplaceAllString(input, "")
}

func trimStringWithNativeTools(input string) string {
	return strings.TrimSpace(input)
}
