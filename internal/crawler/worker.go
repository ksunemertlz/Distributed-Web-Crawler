package crawler

import (
	"sync"

	"context"

	"github.com/ksunemertlz/web-crawler/internal/fetcher"
	"github.com/ksunemertlz/web-crawler/internal/kafka"
	"github.com/ksunemertlz/web-crawler/internal/parser"
	"golang.org/x/time/rate"
)

func Worker(
	ctx context.Context,
	queue chan string,
	store *URLStore,
	wg *sync.WaitGroup,
	limiter *rate.Limiter,
	producer *kafka.Producer) {

	for {
		select {

		case url := <-queue:

			if !store.Add(url) {
				wg.Done()
				continue
			}

			limiter.Wait(ctx)

			html, err := fetcher.Fetch(url)
			if err != nil {
				wg.Done()
				continue
			}

			links := parser.ExtractLinks(html)

			for _, link := range links {
				producer.Send(ctx, link)
			}

			wg.Done()

		case <-ctx.Done():
			return
		}
	}
}
