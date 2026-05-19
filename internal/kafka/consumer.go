package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(broker string, topic string, groupID string) *Consumer {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Start(ctx context.Context) {

	for {

		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return
		}

		fmt.Println("received:", string(msg.Value))
	}
}
