package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/consumer"
)

var (
	receiveKafkaServer string
	receiveTopic       string
	startFrom          string
	receiveGroup       string
)

var ReceiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive messages from Kafka",
	Run: func(_ *cobra.Command, _ []string) {
		// Echo the parameters to the user
		log.Printf("You are receiving from the channel: '%s'\n", receiveTopic)
		log.Printf("You are receiving from the '%s'\n", startFrom)
		log.Printf("You are receiving through the server: '%s'\n", receiveKafkaServer)

		if receiveGroup != "" {
			log.Printf("You are part of the receiving group: '%s'\n", receiveGroup)
		}

		// Create a kafka consumer
		kc, err := consumer.NewKafkaConsumer(receiveKafkaServer, receiveGroup, startFrom)
		if err != nil {
			// return err
			log.Errorf("Failed to create Kafka consumer: %s", err)
		}
		defer kc.Close()

		// Call the ReadMessages function
		if err := kc.ReadMessages(receiveTopic, false); err != nil {
			// return err
			log.Errorf("Error reading messages: %v", err)
		}

	},
}

func SetupReceiveCmd() {
	rootCmd.AddCommand(ReceiveCmd)

	// Define flags for the receive command
	ReceiveCmd.Flags().StringVar(&receiveKafkaServer, "server", "", "Kafka connection string (required)")
	ReceiveCmd.Flags().StringVar(&receiveTopic, "channel", "", "Kafka topic (required)")
	ReceiveCmd.Flags().StringVar(&startFrom, "from", "earliest", "Start consuming from (start|latest)")
	ReceiveCmd.Flags().StringVar(&receiveGroup, "group", "", "Group name (optional)")

	err := ReceiveCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", receiveKafkaServer)
	}

	err = ReceiveCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", receiveTopic)
	}
}
