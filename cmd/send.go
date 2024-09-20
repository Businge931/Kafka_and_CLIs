package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/app"
)

var (
	sendKafkaServer string
	sendTopic       string
	sendGroup       string
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send messages to Kafka",
	Run: func(_ *cobra.Command, _ []string) {
		// Echo the parameters to the user
		log.Printf("You have decided to send to the channel: '%s'\n", sendTopic)
		log.Printf("You are sending through the server: '%s'\n", sendKafkaServer)

		if sendGroup != "" {
			log.Printf("You are sending through the group: '%s'\n", sendGroup)
		}
		// Call the SendMessage function from the app package
		app.SendMessage(sendKafkaServer, sendTopic, sendGroup, false)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Define flags for the send command
	sendCmd.Flags().StringVar(&sendKafkaServer, "server", "", "Kafka connection string (required)")
	sendCmd.Flags().StringVar(&sendTopic, "channel", "", "Kafka topic (required)")
	sendCmd.Flags().StringVar(&sendGroup, "group", "", "Group name (optional)")

	err := sendCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", sendKafkaServer)
	}

	err = sendCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", sendTopic)
	}
}
