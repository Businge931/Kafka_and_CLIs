package broker

import (

	// "github.com/Businge931/Kafka_and_CLIs/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
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

// SendMessage sends a message to the Kafka topic
func (kp *KafkaProducer) SendMessage(topic, message string) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
}

// Close flushes and closes the producer
func (kp *KafkaProducer) Close() {
	kp.producer.Flush(15 * 1000)
	kp.producer.Close()
}
