package wordmaster

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	w      *kafka.Writer
	logger *log.Logger
}

func NewProducer(logger *log.Logger, brokers []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: topic,
	}

	return &Producer{w: w, logger: logger}
}

func (p *Producer) Produce(ctx context.Context, c <-chan *WiktionaryRequest) {
	go func() {
		for {
			select {
			case v, ok := <-c:
				if !ok {
					return
				}
				msg, err := proto.Marshal(v)
				if err != nil {
					p.logger.Errorf("failed to marshal message: %v", err)
					continue
				}

				err = p.w.WriteMessages(context.Background(), kafka.Message{
					Value: msg,
				})
				if err != nil {
					p.logger.Errorf("failed to write message: %v", err)
					continue
				}
			case <-ctx.Done():
				return
			}
		}
	}()

}
