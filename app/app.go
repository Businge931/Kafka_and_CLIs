package app

import (
	"bufio"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

// SendMessage sends messages to the specified Kafka topic
func SendMessage(kafkaServer, topic, _ string, testing bool) {
	// Create Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		return
	}
	defer producer.Close()

	// Delivery report handler for produced messages
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

	reader := bufio.NewReader(os.Stdin)
	message := ""

	for message != "q" {
		if !testing {
			log.Print("Enter message or 'q' to quit: ")
			message, _ = reader.ReadString('\n') // Read the entire line
			message = message[:len(message)-1]   // Remove the newline character
		} else {
			message = "Test message"
		}

		if message == "q" {
			break
		}

		// Produce message to the specified Kafka topic
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
		if err != nil {
			log.Printf("failed to produce message to: %s", topic)
		}

		// Wait for message deliveries before exiting
		producer.Flush(15 * 1000)

		if testing {
			return
		}
	}
}

// ReadMessages consumes messages from the specified Kafka topic
func ReadMessages(kafkaServer, topic, startFrom, group string, testing bool) {
	// Kafka consumer configuration
	offsetReset := "earliest" // Default to start from the beginning
	if startFrom == "latest" {
		offsetReset = "latest"
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          group,
		"auto.offset.reset": offsetReset,
	})
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	// Subscribe to the topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Printf("Failed to subscribe to topic: %s\n", err)
		return
	}

	log.Println("Waiting for messages...")

	running := true
	for running {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))

			if testing {
				running = false
			}
		} else {
			log.Printf("Consumer error: %v\n", err)
			break
		}
	}

	log.Println("Exiting consumer...")
}

// func CreateTopic(kafkaServer string, topicName string, numPartitions int, replicationFactor int) error {
// 	// Create AdminClient instance
// 	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
// 	if err != nil {
// 		return fmt.Errorf("failed to create Admin client: %w", err)
// 	}
// 	defer admin.Close()

// 	// Create a new topic
// 	results, err := admin.CreateTopics(
// 		context.Background(),
// 		[]kafka.TopicSpecification{{
// 			Topic:             topicName,
// 			NumPartitions:     numPartitions,
// 			ReplicationFactor: replicationFactor}},
// 		kafka.SetAdminOperationTimeout(10*time.Second))

// 	if err != nil {
// 		return fmt.Errorf("failed to create topic: %w", err)
// 	}

// 	// Check if the topic was successfully created
// 	for _, result := range results {
// 		if result.Error.Code() != kafka.ErrNoError {
// 			return fmt.Errorf("failed to create topic %s: %v", result.Topic, result.Error.String())
// 		}

// 		log.Printf("Topic %s created successfully\n", result.Topic)
// 	}

// 	return nil
// }
