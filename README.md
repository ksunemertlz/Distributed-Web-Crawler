# Distributed Web Crawler 

### EN
Concurrent web crawler written in Go that demonstrates worker pool concurrency, rate limiting, graceful shutdown and message processing with Kafka.

The crawler downloads web pages, extracts links and publishes them to Kafka for further processing.

### RU
Конкурентный веб-краулер на языке Go, демонстрирующий использование worker pool, ограничение скорости запросов, корректное завершение работы и обработку результатов через Kafka.

Краулер загружает веб-страницы, извлекает ссылки и отправляет их в Kafka для дальнейшей обработки

# Tech Stack
- Go
- Kafka
- goroutines / channels
- rate limiting
- HTML parsing

### Libraries used:

- github.com/segmentio/kafka-go
- golang.org/x/net/html
- golang.org/x/time/rate

# Features

### EN
- Concurrent crawling using goroutines
- Worker pool architecture
- URL queue implemented with channels
- Deduplication of visited URLs
- Request rate limiting
- Graceful shutdown with context
- Kafka integration for distributed processing
- HTML parsing using Go HTML tokenizer

### RU
- Параллельный обход страниц с использованием goroutines
- Архитектура worker pool
- Очередь URL на основе channels
- Защита от повторного обхода страниц (deduplication)
- Ограничение частоты HTTP-запросов (rate limiting)
- Корректное завершение работы (graceful shutdown)
- Интеграция с Kafka для распределённой обработки
- Парсинг HTML-страниц

# Architecture

Seed URLs
   ↓
Crawler Workers
   ↓
Fetch HTML
   ↓
Extract Links
   ↓
Kafka Topic
   ↓
Consumer / Storage

# Project Structure

web-crawler
│
cmd/
   crawler/
      main.go
   consumer/
      main.go
│
internal/
   crawler/
      worker.go
      store.go
   fetcher/
      fetcher.go
   parser/
      parser.go
   kafka/
      producer.go
      consumer.go

# Running the Project
#### 1. Start Kafka
Example with Docker:
docker run -p 9092:9092 apache/kafka
#### 2. Run crawler
go run cmd/crawler/main.go
#### 3. Run consumer
go run cmd/consumer/main.go
Consumer will read URLs from Kafka and print them.