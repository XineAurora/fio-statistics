package message

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type MessageProducer interface {
	Write(ctx context.Context, data []byte) error
	Close() error
}

type KafkaMessageProducer struct {
	writer *kafka.Writer
}

func NewKafkaMessageProducer(writer *kafka.Writer) *KafkaMessageProducer {
	return &KafkaMessageProducer{writer: writer}
}

func (kp *KafkaMessageProducer) Write(ctx context.Context, data []byte) error {
	return kp.writer.WriteMessages(ctx, kafka.Message{Value: data})
}

func (kp *KafkaMessageProducer) Close() error {
	return kp.writer.Close()
}
