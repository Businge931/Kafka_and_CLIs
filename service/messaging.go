package service

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

const MaxRetries = 5

func (svc *Service) SendMessage(topic, message string) error {
	if topic == "" {
		log.Error(models.ErrEmptyTopic)
		return models.ErrEmptyTopic
	}

	if message == "" {
		log.Error(models.ErrEmptyMessage)
		return models.ErrEmptyMessage
	}

	var lastError error

	for attempt := range MaxRetries {
		err := svc.producer.SendMessage(topic, message)
		if err == nil {
			return nil
		}

		lastError = err
		log.Errorf("Failed to send message (attempt %d/%d): %v", attempt+1, MaxRetries, err)
		time.Sleep(time.Duration(500*(1<<attempt)) * time.Millisecond) // Exponential backoff
	}

	return fmt.Errorf("%w: %w", models.ErrMaxRetry, lastError)
}

func (svc *Service) ReadMessages(topic string) error {
	if topic == "" {
		log.Error(models.ErrEmptyTopic)
		return models.ErrEmptyTopic
	}

	err := svc.consumer.ReadMessages(topic)
	if err != nil {
		log.Errorf("Failed to receive messages from topic %s: %v", topic, err)
		return fmt.Errorf("error receiving messages from topic %s: %w", topic, err)
	}

	return nil
}
