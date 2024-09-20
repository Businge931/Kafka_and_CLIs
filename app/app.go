package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// SendMessage sends messages to the specified Kafka topic
func SendMessage(kafkaServer string, topic string, group string, testing bool) {
	// Create Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	message := ""
	for message != "q" {
		if !testing {
			fmt.Print("Enter message or 'q' to quit: ")
			message, _ = reader.ReadString('\n') // Read the entire line
			message = message[:len(message)-1]   // Remove the newline character
		} else {
			message = "Test message"
		}

		if message == "q" {
			break
		}

		// Produce message to the specified Kafka topic
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
			Value:          []byte(message),
		}, nil)

		// Wait for message deliveries before exiting
		p.Flush(15 * 1000)
		if testing {
			return
		}
	}
}

// ReadMessages consumes messages from the specified Kafka topic
func ReadMessages(kafkaServer string, topic string, startFrom string, group string, testing bool) {
	// Kafka consumer configuration
	offsetReset := "earliest" // Default to start from the beginning
	if startFrom == "latest" {
		offsetReset = "latest"
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          group,
		"auto.offset.reset": offsetReset,
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer c.Close()

	// Subscribe to the topic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s\n", err)
		return
	}

	fmt.Println("Waiting for messages...")

	running := true
	for running {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
			if testing {
				running = false
			}
		} else {
			fmt.Printf("Consumer error: %v\n", err)
			break
		}
	}

	fmt.Println("Exiting consumer...")
}

func CreateTopic(kafkaServer string, topicName string, numPartitions int, replicationFactor int) error {
	// Create AdminClient instance
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		return fmt.Errorf("Failed to create Admin client: %s", err)
	}
	defer admin.Close()

	// Create a new topic
	results, err := admin.CreateTopics(
		context.Background(),
		[]kafka.TopicSpecification{{
			Topic:             topicName,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor}},
		kafka.SetAdminOperationTimeout(10*time.Second))

	if err != nil {
		return fmt.Errorf("Failed to create topic: %v", err)
	}

	// Check if the topic was successfully created
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			return fmt.Errorf("Failed to create topic %s: %v", result.Topic, result.Error.String())
		}
		fmt.Printf("Topic %s created successfully\n", result.Topic)
	}

	return nil
}
