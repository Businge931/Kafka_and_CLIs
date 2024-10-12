/*
Copyright © 2024 BUSINGE BISANGA <busingebisanga99@gmail.com>
*/
package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/Businge931/Kafka_and_CLIs/broker"
	"github.com/Businge931/Kafka_and_CLIs/cmd"
	"github.com/Businge931/Kafka_and_CLIs/service"
)

func main() {
	// Get Kafka server address from environment or configuration
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found or failed to load")
	}
	kafkaServer := os.Getenv("KAFKA_SERVER")
	if kafkaServer == "" {
		log.Fatal("KAFKA_SERVER environment variable is not set")
	} else {
		log.Infof("Kafka server address: %s", kafkaServer) 
	}

	// Initialize the Kafka producer
	kafkaProducer, err := broker.NewProducer(kafkaServer)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	// Initialize the Kafka consumer
	kafkaConsumer, err := broker.NewConsumer(kafkaServer, "tests", "earliest", false)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	// Wrap the Kafka producer in a generic service layer
	messageService := service.NewMessageService(kafkaProducer, kafkaConsumer)

	// Call the setup functions to initialize commands
	cmd.SetupSendCmd(messageService)
	cmd.SetupReceiveCmd(messageService)
	cmd.SetupRootCmd()

	// Execute the root command
	cmd.Execute()
}
