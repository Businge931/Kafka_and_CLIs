package cmd

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/broker"
	"github.com/Businge931/Kafka_and_CLIs/service"
)

var (
	sendKafkaServer string
	sendTopic       string
	sendGroup       string
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
		kafkaProducer, err := broker.NewKafkaProducer(sendKafkaServer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create Kafka producer: %v\n", err)
			os.Exit(1)
		}
		defer kafkaProducer.Close()

		// Initialize the generic producer with Kafka-specific logic
		genericProducer := broker.NewProducer(kafkaProducer.SendMessage, kafkaProducer.Close)

		// Initialize the service layer
		svc := service.New(genericProducer, nil)

		// Capture user input and send messages using the service
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter message (or 'q' to quit): ")
			message, _ := reader.ReadString('\n')

			// Remove newline and quit if the user types 'q'
			message = message[:len(message)-1]
			if message == "q" {
				break
			}
			// Send the message using the service
			if err := svc.SendMessage(sendTopic, message); err != nil {
				log.Fatalf("Failed to send message using producer: %v", err)
			}
		}

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
