package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/XineAurora/fio-statistics/intrernal/entities"
	_ "github.com/joho/godotenv/autoload"
	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_FIO_TOPIC"), 0)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 5))

	fio := entities.FIO{Name: "Joseph", Surname: "Jostar", Patronymic: "George"}

	data, err := json.Marshal(fio)
	if err != nil {
		log.Fatal(err)
	}

	conn.WriteMessages(kafka.Message{Value: data})
}
