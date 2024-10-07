package broker

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

// NewKafkaConsumer creates a new KafkaConsumer
func NewKafkaConsumer(kafkaServer, groupPrefix, startFrom string, dynamicGroup bool) (*KafkaConsumer, error) {
	config := configureConsumerOptions(startFrom)

	// Dynamically create unique group ID if dynamicGroup is true
	group := groupPrefix
	if dynamicGroup {
		group = fmt.Sprintf("%s-%d", groupPrefix, time.Now().UnixNano()) // Unique group ID
	}

	consumer, err := createConsumer(kafkaServer, group, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %s", err)
	}

	return &KafkaConsumer{consumer: consumer}, nil
}

// ReadMessages reads messages from the Kafka topic
func (kc *KafkaConsumer) ReadMessages(topic string) error {
	if err := kc.consumer.SubscribeTopics([]string{topic}, nil); err != nil {
		return fmt.Errorf("failed to subscribe to topic: %s", err)
	}

	log.Println("Waiting for messages...")
	log.Println("Subscribed to topic:", topic)
	return readMessages(kc.consumer)
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

func readMessages(consumer *kafka.Consumer) error {
	for {
		msg, err := consumer.ReadMessage(-1) // The parameter -1 specifies that the method should block the process and wait for a message indefinitely until a message is received.
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))
		} else {
			return fmt.Errorf("consumer error: %v", err)
		}
	}
}
