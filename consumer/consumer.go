package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func ReadMessages(kafkaServer, topic, startFrom, group string, testing bool) {
	config := configureConsumerOptions(startFrom)

	consumer, err := createConsumer(kafkaServer, group, config)
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	if err := subscribeToTopic(consumer, topic); err != nil {
		log.Printf("Failed to subscribe to topic: %s\n", err)
		return
	}

	log.Println("Waiting for messages...")

	if err := readMessages(consumer, testing); err != nil {
		log.Printf("Error reading messages: %v\n", err)
	}

	log.Println("Exiting consumer...")
}

func configureConsumerOptions(startFrom string) kafka.ConfigMap {
	offsetReset := "earliest"
	if startFrom == "latest" {
		offsetReset = "latest"
	}
	return kafka.ConfigMap{"auto.offset.reset": offsetReset}
}

func createConsumer(kafkaServer, group string, config kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          group,
		"auto.offset.reset": config["auto.offset.reset"],
	})
}

func subscribeToTopic(consumer *kafka.Consumer, topic string) error {
	return consumer.SubscribeTopics([]string{topic}, nil)
}

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
