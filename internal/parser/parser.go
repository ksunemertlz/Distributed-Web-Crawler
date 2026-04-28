package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlStr string) []string {
	var links []string

	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return links
	}

	var walk func(*html.Node)

	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {

				if attr.Key == "href" {
					links = append(links, attr.Val)
				}

			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	return links
}
