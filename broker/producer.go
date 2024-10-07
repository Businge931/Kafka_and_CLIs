package broker

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(kafkaServer string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		return nil, err
	}

	kp := &KafkaProducer{producer: producer}
	kp.startDeliveryReportHandler()

	return kp, nil
}

// startDeliveryReportHandler starts a goroutine to handle delivery reports
func (kp *KafkaProducer) startDeliveryReportHandler() {
	go func() {
		for e := range kp.producer.Events() {
			if ev, ok := e.(*kafka.Message); ok {
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\\n", ev.TopicPartition)
				}
			}
		}
	}()
}

// SendMessage sends a message to the Kafka topic with retries
func (kp *KafkaProducer) SendMessage(topic, message string) error {
	const maxRetries = 5
	const baseRetryDelay = 500

	var lastError error

	for attempt := range [maxRetries]int{} {
		// Attempt to send the message
		err := kp.sendWithDeliveryReport(topic, message)
		if err == nil {
			return nil
		}

		// Log the failure and retry
		lastError = err
		log.Errorf("Failed to send message (attempt %d/%d): %v. Retrying...", attempt+1, maxRetries, err)

		// Exponential backoff for retries
		time.Sleep(time.Duration(baseRetryDelay*(1<<attempt)) * time.Millisecond)
	}

	return fmt.Errorf("%w: %w", models.ErrMaxRetry, lastError)
}

// Close flushes and closes the producer
func (kp *KafkaProducer) Close() {
	kp.producer.Flush(15 * 1000)
	kp.producer.Close()
}

// sendWithDeliveryReport is a helper that sends the message and checks for delivery report errors
func (kp *KafkaProducer) sendWithDeliveryReport(topic, message string) error {
	deliveryChan := make(chan kafka.Event)

	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return err
	}

	// Wait for delivery report
	e := <-deliveryChan

	msg, ok := e.(*kafka.Message)
	if !ok {
		close(deliveryChan)
		return fmt.Errorf("%w: %T", models.ErrEventType, e)
	}

	close(deliveryChan)

	// Check if the message was delivered successfully
	if msg.TopicPartition.Error != nil {
		return msg.TopicPartition.Error
	}

	log.Printf("Message delivered to %v\n", msg.TopicPartition)
	return nil
}
