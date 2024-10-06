package cmd

import (
	// "context"
	// "fmt"
	// "os"

	// "time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/broker"
	"github.com/Businge931/Kafka_and_CLIs/service"
)

var (
	receiveServer string
	receiveTopic  string
	startFrom     string
	receiveGroup  string
	// testing       bool
)

// ReceiveCmd defines the Cobra command for consuming messages
var ReceiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive messages from Kafka",
	Run: func(_ *cobra.Command, _ []string) {
		// Echo the parameters to the user
		log.Printf("You are receiving from the channel: '%s'\n", receiveTopic)
		log.Printf("You are receiving from the '%s'\n", startFrom)
		log.Printf("You are receiving through the server: '%s'\n", receiveServer)

		if receiveGroup != "" {
			log.Printf("You are part of the receiving group: '%s'\n", receiveGroup)
		}

		// Set up a context with a timeout for receiving messages
		// ctx, cancel := context.WithCancel(context.Background())
		// defer cancel()

		// Initialize Kafka consumer
		kafkaConsumer, err := broker.NewKafkaConsumer(receiveServer, receiveGroup, startFrom)
		if err != nil {
			log.Fatalf("Failed to create Kafka consumer: %s\n", err)

		}
		defer kafkaConsumer.Close()

		// Initialize the generic consumer with Kafka-specific logic
		genericConsumer := broker.NewConsumer(kafkaConsumer.ReadMessages, kafkaConsumer.Close)

		// Initialize the service layer
		svc := service.New(nil, genericConsumer)

		// Receive messages using the service
		if err := svc.ReceiveMessages(receiveTopic); err != nil {
			log.Fatalf("Failed to receive messages using consumer: %s", err)
		}

	},
}

// SetupReceiveCmd sets up the flags and adds the command to the root
func SetupReceiveCmd() {
	rootCmd.AddCommand(ReceiveCmd)

	// Define flags for the receive command
	ReceiveCmd.Flags().StringVar(&receiveServer, "server", "", "Kafka connection string (required)")
	ReceiveCmd.Flags().StringVar(&receiveTopic, "channel", "", "Kafka topic (required)")
	ReceiveCmd.Flags().StringVar(&startFrom, "from", "earliest", "Start consuming from (start|latest)")
	ReceiveCmd.Flags().StringVar(&receiveGroup, "group", "", "Group name (optional)")

	err := ReceiveCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", receiveServer)
	}

	err = ReceiveCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", receiveTopic)
	}

}
