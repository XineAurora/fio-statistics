package message

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type MessageConsumer interface {
	Read(ctx context.Context) ([]byte, error)
	Close() error
}

type KafkaMessageConsumer struct {
	reader *kafka.Reader
}

func NewKafkaMessageConsumer(reader *kafka.Reader) *KafkaMessageConsumer {
	return &KafkaMessageConsumer{reader: reader}
}

func (mc *KafkaMessageConsumer) Read(ctx context.Context) ([]byte, error) {
	msg, err := mc.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return msg.Value, nil
}

func (mc *KafkaMessageConsumer) Close() error {
	return mc.reader.Close()
}
