package message

import (
	"context"
	"encoding/json"
	"log"

	"github.com/XineAurora/fio-statistics/intrernal/database"
	"github.com/XineAurora/fio-statistics/intrernal/enricher"
	"github.com/XineAurora/fio-statistics/intrernal/entities"
)

type MessageHandler struct {
	consumer    MessageConsumer
	producer    MessageProducer
	repo        database.FIORepository
	enricher    enricher.Enricher
	quitChan    chan bool
	messageChan chan []byte
	errorChan   chan error
}

func NewMessageHandler(consumer MessageConsumer, producer MessageProducer,
	repo database.FIORepository, enricher enricher.Enricher) *MessageHandler {
	return &MessageHandler{
		consumer:    consumer,
		producer:    producer,
		repo:        repo,
		enricher:    enricher,
		quitChan:    make(chan bool),
		messageChan: make(chan []byte),
		errorChan:   make(chan error),
	}
}

func (m *MessageHandler) Start() error {
	ctx, cancelReader := context.WithCancel(context.Background())
	defer cancelReader()
	go func() {
		for {
			data, err := m.consumer.Read(ctx)
			select {
			case <-ctx.Done():
				return
			default:
			}
			if err != nil {
				m.errorChan <- err
			}
			m.messageChan <- data
		}
	}()

	for {
		select {
		case <-m.quitChan:
			return nil
		case msgData := <-m.messageChan:
			go m.tryCreateFIO(msgData)
		case err := <-m.errorChan:
			m.Stop()
			return err
		}
	}
}

func (m *MessageHandler) Stop() error {
	m.quitChan <- true

	if err := m.consumer.Close(); err != nil {
		return err
	}
	if err := m.producer.Close(); err != nil {
		return err
	}
	return nil
}

func (m *MessageHandler) tryCreateFIO(msgData []byte) {
	fioBasic, err := entities.NewFIO(msgData)
	if err != nil {
		failed, err := json.Marshal(entities.FIOFailed{RawData: string(msgData), ErrorMessage: err.Error()})
		if err != nil {
			log.Printf("marshal error: %s\n", err.Error())
			m.errorChan <- err
		}
		err = m.producer.Write(context.Background(), failed)
		if err != nil {
			log.Printf("producer write error: %s\n", err)
			m.errorChan <- err
		}

	}
	fio, err := fioBasic.EnrichFIO(m.enricher)
	if err != nil {
		log.Printf("enrichment error: %s\n", err)
		m.errorChan <- err
	}
	_, err = m.repo.CreateFIO(fio)
	if err != nil {
		log.Printf("fio saving error: %s\n", err)
		m.errorChan <- err
	}
}
