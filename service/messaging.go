package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

func (svc *Service) SendMessage(topic, message string) error {
	const maxRetries = 5

	if topic == "" {
		log.Error(models.ErrEmptyTopic)
		return models.ErrEmptyTopic
	}

	if message == "" {
		log.Error(models.ErrEmptyMessage)
		return models.ErrEmptyMessage
	}

	var lastError error
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := svc.producer.SendMessage(topic, message)
		if err == nil {
			return nil
		}

		lastError = err
		log.Errorf("Failed to send message (attempt %d/%d): %v", attempt+1, maxRetries, err)
		time.Sleep(time.Duration(500*(1<<attempt)) * time.Millisecond) // Exponential backoff
	}

	return fmt.Errorf("%w: %v", models.ErrMaxRetry, lastError)
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

// SendHandler sends messages using the provided Service.
func SendHandler(svc MessageService, topic string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("Enter message (or 'q' to quit): ")
		message, _ := reader.ReadString('\n')

		// Remove the newline character
		message = strings.TrimSpace(message)
		if message == "q" {
			break
		}

		// Send the message using the service
		if err := svc.SendMessage(topic, message); err != nil {
			log.Errorf("Failed to send message: %v", err)
		}
	}
}