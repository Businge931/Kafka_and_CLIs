package producer

import (
	"bufio"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

// MessageProducer defines the interface for producing messages
type MessageProducer interface {
	SendMessage(topic, message string) error
	Close()
}

// KafkaProducer implements the MessageProducer interface
type KafkaProducer struct {
	producer *kafka.Producer
}

// Define a variable that holds the producer creation function
// var KafkaProducerFactory = NewKafkaProducer

// NewKafkaProducer creates a new KafkaProducer
func NewKafkaProducer(kafkaServer string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		return nil, err
	}

	kp := &KafkaProducer{producer: producer}
	kp.startDeliveryReportHandler()
	return kp, nil
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

// getMessageInput is used to get input from the user or return a test message
func getMessageInput(testing bool) string {
	if testing {
		return "Test message"
	}
	reader := bufio.NewReader(os.Stdin)
	log.Print("Enter message or 'q' to quit: ")
	message, _ := reader.ReadString('\n')
	return message[:len(message)-1]
}

// Producer logic to send messages
func ProduceMessages(kp MessageProducer, topic string, testing bool) {
	for {
		message := getMessageInput(testing)
		if message == "q" {
			break
		}

		if err := kp.SendMessage(topic, message); err != nil {
			log.Printf("Failed to produce message: %s", err)
		}

		if testing {
			return
		}
	}
}
