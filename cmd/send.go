package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		// Echo the parameters to the user
		fmt.Printf("You have decided to send to the channel: '%s'\n", sendTopic)
		fmt.Printf("You are sending through the server: '%s'\n", sendKafkaServer)
		if sendGroup != "" {
			fmt.Printf("You are sending through the group: '%s'\n", sendGroup)
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

	sendCmd.MarkFlagRequired("server")
	sendCmd.MarkFlagRequired("channel")
}
