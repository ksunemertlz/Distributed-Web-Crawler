package main

import (
	"sync"

	"github.com/ksunemertlz/web-crawler/internal/crawler"
	"github.com/ksunemertlz/web-crawler/internal/queue"
)

func main() {

	urlQueue := queue.NewQueue()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go crawler.Worker(urlQueue, &wg)
	}

	startURL := "https://example.com"

	wg.Add(1)
	urlQueue <- startURL

	wg.Wait()
}
