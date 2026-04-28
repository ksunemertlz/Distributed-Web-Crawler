package main

import (
	"fmt"

	"github.com/ksunemertlz/web-crawler/internal/fetcher"
	"github.com/ksunemertlz/web-crawler/internal/parser"
)

func main() {
	html, err := fetcher.Fetch("https://example.com")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	links := parser.ExtractLinks(html)

	for _, l := range links {
		fmt.Println(l)
	}
}
