package service

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

func (svc *Service) SendMessage(topic, message string) error {
	if topic == "" {
		log.Error(models.ErrEmptyTopic)
		return models.ErrEmptyTopic
	}

	if message == "" {
		log.Error(models.ErrEmptyMessage)
		return models.ErrEmptyMessage
	}

	err := svc.producer.SendMessage(topic, message)
	if err != nil {
		log.Errorf("Failed to send message to topic %s: %v", topic, err)
		return fmt.Errorf("error sending message to topic %s: %w", topic, err)
	}

	return nil
}

func (svc *Service) ReceiveMessages(topic string) error {
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
