package intrernal

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageConsumer struct {
	reader *kafka.Reader
}

func NewKafkaMessageConsumer(reader *kafka.Reader) KafkaMessageConsumer {
	return KafkaMessageConsumer{reader: reader}
}

func (mc *KafkaMessageConsumer) Read() ([]byte, error) {
	msg, err := mc.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	return msg.Value, nil
}
