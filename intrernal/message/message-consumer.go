package message

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// const (
// 	KAFKA_HOST             = "KAFKA_HOST"
// 	KAFKA_FIO_TOPIC        = "KAFKA_FIO_TOPIC"
// 	KAFKA_FIO_FAILED_TOPIC = "KAFKA_FIO_FAILED_TOPIC"
// )

type MessageConsumer interface {
	Read(ctx context.Context) ([]byte, error)
	Close() error
}

type KafkaMessageConsumer struct {
	reader *kafka.Reader
}

// r := kafka.NewReader(kafka.ReaderConfig{
// 	Brokers:  []string{os.Getenv(KAFKA_HOST)},
// 	Topic:    os.Getenv(KAFKA_FIO_TOPIC),
// 	GroupID:  "test-group",
// 	MaxBytes: 1e6,
// })
// defer func() {
// 	if err := r.Close(); err != nil {
// 		log.Fatal("failed to close reader:", err)
// 	}
// }()

func NewKafkaMessageConsumer(reader *kafka.Reader) KafkaMessageConsumer {
	return KafkaMessageConsumer{reader: reader}
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
