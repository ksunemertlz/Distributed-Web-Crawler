package crawler

import (
	"fmt"
	"sync"

	"github.com/ksunemertlz/web-crawler/internal/fetcher"
	"github.com/ksunemertlz/web-crawler/internal/parser"
)

func Worker(queue chan string, wg *sync.WaitGroup) {

	for url := range queue {

		fmt.Println("Crawling:", url)

		html, err := fetcher.Fetch(url)
		if err != nil {
			wg.Done()
			continue
		}

		links := parser.ExtractLinks(html)

		for _, link := range links {
			queue <- link
			wg.Add(1)
		}

		wg.Done()
	}
}
