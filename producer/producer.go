package producer

import (
	"bufio"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)


func SendMessage(kafkaServer, topic, _ string, testing bool) {
    producer, err := createProducer(kafkaServer)
    if err != nil {
        log.Printf("Failed to create producer: %s\n", err)
        return
    }
    defer producer.Close()

    startDeliveryReportHandler(producer)

    for {
        message := getMessageInput(testing)
        if message == "q" {
            break
        }

        if err := _sendMessage(producer, topic, message); err != nil {
            log.Printf("Failed to produce message: %s", err)
        }

        flushProducer(producer)

        if testing {
            return
        }
    }
}

func createProducer(kafkaServer string) (*kafka.Producer, error) {
    return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
}

func startDeliveryReportHandler(producer *kafka.Producer) {
    go func() {
        for e := range producer.Events() {
            if ev, ok := e.(*kafka.Message); ok {
                if ev.TopicPartition.Error != nil {
                    log.Printf("Delivery failed: %v\n", ev.TopicPartition)
                } else {
                    log.Printf("Delivered message to %v\n", ev.TopicPartition)
                }
            }
        }
    }()
}

func getMessageInput(testing bool) string {
    if testing {
        return "Test message"
    }
    reader := bufio.NewReader(os.Stdin)
    log.Print("Enter message or 'q' to quit: ")
    message, _ := reader.ReadString('\n')
    return message[:len(message)-1]
}

func _sendMessage(producer *kafka.Producer, topic, message string) error {
    return producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)
}

func flushProducer(producer *kafka.Producer) {
    producer.Flush(15 * 1000)
}
