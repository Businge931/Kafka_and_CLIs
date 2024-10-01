package kafka

import (
	"context"
	"fmt"

	"github.com/Businge931/Kafka_and_CLIs/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

// Consumer implements the MessageConsumer interface from the service layer
type Consumer struct {
	consumer *kafka.Consumer
}

// NewKafkaConsumer creates and returns a new Consumer instance
func NewConsumer(kafkaServer, group, startFrom string) (*Consumer, error) {
	config := configureConsumerOptions(startFrom)
	consumer, err := createConsumer(kafkaServer, group, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %s", err)
	}

	return &Consumer{consumer: consumer}, nil
}


// ConsumeMessages reads messages from the Kafka topic (implements the MessageConsumer interface)
func (kc *Consumer) ConsumeMessages(ctx context.Context) ([]models.Message, error) {
	var messages []models.Message

	// Read messages from the Kafka topic
	for {
		select {
		case <-ctx.Done():
			return messages, ctx.Err()
		default:
			msg, err := kc.consumer.ReadMessage(-1)
			if err == nil {
				// Convert Kafka message to models.Message and append it to the slice
				messages = append(messages, models.Message{
					Channel:  *msg.TopicPartition.Topic,
					Server:   "", // Server info may not be available from the message directly
					Group:    "", // Group info might be handled differently
					Payload:  string(msg.Value),
					Metadata: extractHeaders(msg.Headers),
				})
			} else {
				log.Errorf("Consumer error: %v", err)
				return messages, err
			}
		}
	}
}

// Close closes the Kafka consumer
func (kc *Consumer) Close() {
	kc.consumer.Close()
}


// configureConsumerOptions configures Kafka consumer options based on the start offset
func configureConsumerOptions(startFrom string) kafka.ConfigMap {
	offsetReset := "earliest"
	if startFrom == "latest" {
		offsetReset = "latest"
	}
	return kafka.ConfigMap{"auto.offset.reset": offsetReset}
}

// createConsumer creates a new Kafka consumer with the specified configuration
func createConsumer(kafkaServer, group string, config kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          group,
		"auto.offset.reset": config["auto.offset.reset"],
	})
}

// extractHeaders converts Kafka headers to a map for models.Message metadata
func extractHeaders(headers []kafka.Header) map[string]string {
	metadata := make(map[string]string)
	for _, header := range headers {
		metadata[header.Key] = string(header.Value)
	}
	return metadata
}