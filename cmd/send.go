package cmd

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/kafka"
	"github.com/Businge931/Kafka_and_CLIs/models"
	"github.com/Businge931/Kafka_and_CLIs/service"
)

var (
	sendKafkaServer string
	sendTopic       string
	sendGroup       string
	payload         string
	metadata        map[string]string
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send messages to Kafka",
	Run: func(_ *cobra.Command, _ []string) {
		// Echo the parameters to the user
		log.Printf("You have decided to send to the channel: '%s'\n", sendTopic)
		log.Printf("You are sending through the server: '%s'\n", sendKafkaServer)

		if sendGroup != "" {
			log.Printf("You are sending through the group: '%s'\n", sendGroup)
		}
		// Create a new Kafka producer
		producer, err := kafka.NewProducer(sendKafkaServer)
		if err != nil {
			log.Fatalf("Failed to create Kafka producer: %v", err)
		}
		defer producer.Close()

		// Initialize the service with the Kafka producer
		msgService := service.New(producer, nil)

		// Create a message based on input flags
		message := models.Message{
			Channel:  sendTopic,
			Server:   sendKafkaServer,
			Group:    sendGroup,
			Payload:  payload,
			Metadata: metadata,
		}

		// Send the message through the service layer
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := msgService.SendMessage(ctx, message); err != nil {
			log.Fatalf("Failed to start kafka producer: %v", err)
		}

		log.Println("Message sent successfully.")
	},
}

func SetupSendCmd() {
	rootCmd.AddCommand(SendCmd)

	// Define flags for the send command
	SendCmd.Flags().StringVar(&sendKafkaServer, "server", "", "Kafka connection string (required)")
	SendCmd.Flags().StringVar(&sendTopic, "channel", "", "Kafka topic (required)")
	SendCmd.Flags().StringVar(&sendGroup, "group", "", "Group name (optional)")

	err := SendCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", sendKafkaServer)
	}

	err = SendCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", sendTopic)
	}
}
