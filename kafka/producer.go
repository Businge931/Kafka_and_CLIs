package kafka

import (
	"context"
	"time"

	"github.com/Businge931/Kafka_and_CLIs/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

type Producer struct {
	producer *kafka.Producer
}

// NewKafkaProducer creates and returns a new Producer instance
func NewProducer(kafkaServer string) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		return nil, err
	}

	kp := &Producer{producer: producer}
	kp.startDeliveryReportHandler()
	return kp, nil
}

func (kp *Producer) ProduceMessage(
	ctx context.Context,
	 message models.Message) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &message.Channel, Partition: kafka.PartitionAny},
		Value:          []byte(message.Payload),
		Headers:        convertHeaders(message.Metadata),
		Timestamp:      time.Now(),
	}

	// Produces the message
	// return kp.producer.Produce(kafkaMsg, nil);
	if err := kp.producer.Produce(kafkaMsg, nil); err != nil {
		return err
	}


	// Optional: wait for message delivery (or you can handle it asynchronously)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second): // Simulate waiting for a delivery report (non-blocking)
		return nil
	}
	
}

// Close flushes and closes the Kafka producer
func (kp *Producer) Close() {
	kp.producer.Flush(15 * 1000)
	kp.producer.Close()
}

// startDeliveryReportHandler handles delivery reports from Kafka asynchronously
func (kp *Producer) startDeliveryReportHandler() {
	go func() {
		for e := range kp.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					log.Infof("Message delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

// Helper function to convert headers from map[string]string to kafka.Headers format
func convertHeaders(headers map[string]string) []kafka.Header {
	kafkaHeaders := make([]kafka.Header, 0, len(headers))
	for key, value := range headers {
		kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: key, Value: []byte(value)})
	}
	return kafkaHeaders
}
