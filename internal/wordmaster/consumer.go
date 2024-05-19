package wordmaster

import (
	"context"
	"errors"
	"github.com/charmbracelet/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"io"
	"time"
)

type Consumer struct {
	r      *kafka.Reader
	logger *log.Logger
}

func NewConsumer(logger *log.Logger, brokers []string, topic string, maxWait time.Duration) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "aboba",
		MaxWait: maxWait,
	})

	return &Consumer{r: r, logger: logger}
}

func (c *Consumer) Close() error {
	return c.r.Close()
}

func (c *Consumer) Consume(ctx context.Context) <-chan *WiktionaryResponse {
	responses := make(chan *WiktionaryResponse)

	go func() {
		defer close(responses)
		for {
			m, err := c.r.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, io.EOF) {
					c.logger.Info("cancelled", "ctx", err)
					return
				}

				c.logger.Errorf("failed to read message: %v", err)
				continue
			}

			r := &WiktionaryResponse{}
			if err := proto.Unmarshal(m.Value, r); err != nil {
				c.logger.Errorf("failed to unmarshal message: %v", err)
				continue
			}

			select {
			case responses <- r:
			case <-ctx.Done():
				return
			}
		}
	}()

	return responses
}
