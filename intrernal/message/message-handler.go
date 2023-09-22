package message

import (
	"context"
	"encoding/json"
	"fmt"

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
			if err != nil {
				if err == ctx.Err() {
					return
				}
				m.errorChan <- fmt.Errorf("consumer read error: %s", err)
			}
			// try parse data, if has error discard it
			fioBasic, err := entities.NewFIO(data)

			if err != nil {
				failed, _ := json.Marshal(entities.FIOFailed{RawData: string(data), ErrorMessage: err.Error()})
				err = m.producer.Write(ctx, failed)
				if err != nil {
					if err == ctx.Err() {
						return
					}
					m.errorChan <- fmt.Errorf("producer write error: %s", err)
				}
			}
			// enrich fio
			fio, err := fioBasic.EnrichFIO(m.enricher)
			if err != nil {
				m.errorChan <- fmt.Errorf("enrichment error: %s", err)
			}

			// put in repo
			_, err = m.repo.CreateFIO(fio)
			if err != nil {
				m.errorChan <- fmt.Errorf("fio saving error: %s", err)
			}
		}
	}()

	select {
	case <-m.quitChan:
		return nil
	case err := <-m.errorChan:
		return err
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
