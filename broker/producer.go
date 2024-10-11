package broker

import (
	"fmt"
	// "time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/service"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

var _ service.MessageProducer = (*KafkaProducer)(nil)

func NewProducer(kafkaServer string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
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

func (kp *KafkaProducer) SendMessage(topic, message string) error {
	deliveryChan := make(chan kafka.Event)

	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return fmt.Errorf("produce error: %w", err)
	}

	e := <-deliveryChan
	close(deliveryChan)

	if msg, ok := e.(*kafka.Message); ok && msg.TopicPartition.Error != nil {
		return fmt.Errorf("delivery error: %w", msg.TopicPartition.Error)
	}

	return nil
}

// Close flushes and closes the producer
func (kp *KafkaProducer) Close() {
	log.Info("Flushing pending messages...")
	kp.producer.Flush(15 * 1000) // 15 seconds timeout
	kp.producer.Close()
	log.Info("Kafka producer closed.")
}
