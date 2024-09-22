package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

// MessageConsumer defines the interface for consuming messages
type MessageConsumer interface {
	ReadMessages(topic string, testing bool) error
	Close()
}

// KafkaConsumer implements the MessageConsumer interface
type KafkaConsumer struct {
	consumer *kafka.Consumer
}

// NewKafkaConsumer creates a new KafkaConsumer
func NewKafkaConsumer(kafkaServer, group, startFrom string) (*KafkaConsumer, error) {
	config := configureConsumerOptions(startFrom)
	consumer, err := createConsumer(kafkaServer, group, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %s", err)
	}

	return &KafkaConsumer{consumer: consumer}, nil
}

// ReadMessages reads messages from the Kafka topic
func (kc *KafkaConsumer) ReadMessages(topic string, testing bool) error {
	if err := kc.consumer.SubscribeTopics([]string{topic}, nil); err != nil {
		return fmt.Errorf("failed to subscribe to topic: %s", err)
	}

	log.Println("Waiting for messages...")
	return readMessages(kc.consumer, testing)
}

// Close closes the Kafka consumer
func (kc *KafkaConsumer) Close() {
	kc.consumer.Close()
}

// configureConsumerOptions configures Kafka consumer options
func configureConsumerOptions(startFrom string) kafka.ConfigMap {
	offsetReset := "earliest"
	if startFrom == "latest" {
		offsetReset = "latest"
	}
	return kafka.ConfigMap{"auto.offset.reset": offsetReset}
}

// createConsumer creates a new Kafka consumer
func createConsumer(kafkaServer, group string, config kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          group,
		"auto.offset.reset": config["auto.offset.reset"],
	})
}

// readMessages reads and processes messages from the Kafka topic
func readMessages(consumer *kafka.Consumer, testing bool) error {
	running := true
	for running {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))
			if testing {
				running = false
			}
		} else {
			return fmt.Errorf("consumer error: %v", err)
		}
	}
	return nil
}
