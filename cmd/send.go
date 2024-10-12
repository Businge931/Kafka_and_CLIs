package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/service"
)

var (
	sendKafkaServer string
	sendTopic       string
	sendGroup       string
	// sendService will hold the reference to the injected MessageService.
	sendService service.MessageService
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

		// Delegate message sending to the injected service
		service.SendHandler(sendService, sendTopic)
	},
}

func SetupSendCmd(svc service.MessageService) {
	sendService = svc // Inject the MessageService instance

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
