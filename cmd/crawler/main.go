package main

import (
	"sync"

	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ksunemertlz/web-crawler/internal/crawler"
	"github.com/ksunemertlz/web-crawler/internal/kafka"
	"github.com/ksunemertlz/web-crawler/internal/queue"
	"golang.org/x/time/rate"
)

func main() {

	urlQueue := queue.NewQueue()
	store := crawler.NewURLStore()

	limiter := rate.NewLimiter(2, 5)

	producer := kafka.NewProducer(
		"localhost:9092",
		"crawler-links",
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()

	for i := 0; i < 5; i++ {
		go crawler.Worker(ctx, urlQueue, store, &wg, limiter, producer)
	}

	startURL := "https://example.com"

	wg.Add(1)
	urlQueue <- startURL

	wg.Wait()
}
