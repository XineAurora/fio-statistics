package intrernal

import "github.com/segmentio/kafka-go"

type KafkaMessageConsumer struct {
	conn *kafka.Conn
}

func NewKafkaMessageConsumer(conn *kafka.Conn) KafkaMessageConsumer {
	return KafkaMessageConsumer{conn: conn}
}
