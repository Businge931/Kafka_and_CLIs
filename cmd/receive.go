package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	
	"github.com/Businge931/Kafka_and_CLIs/app"
)

var (
	receiveKafkaServer string
	receiveTopic       string
	startFrom          string
	receiveGroup       string
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive messages from Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		// Echo the parameters to the user
		fmt.Printf("You are receiving from the channel: '%s'\n", receiveTopic)
		fmt.Printf("You are receiving from the '%s'\n", startFrom)
		fmt.Printf("You are receiving through the server: '%s'\n", receiveKafkaServer)
		if receiveGroup != "" {
			fmt.Printf("You are part of the receiving group: '%s'\n", receiveGroup)
		}
		// Call the ReadMessages function from the app package
		app.ReadMessages(receiveKafkaServer, receiveTopic, startFrom, receiveGroup, false)
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)

	// Define flags for the receive command
	receiveCmd.Flags().StringVar(&receiveKafkaServer, "server", "", "Kafka connection string (required)")
	receiveCmd.Flags().StringVar(&receiveTopic, "channel", "", "Kafka topic (required)")
	receiveCmd.Flags().StringVar(&startFrom, "from", "earliest", "Start consuming from (start|latest)")
	receiveCmd.Flags().StringVar(&receiveGroup, "group", "", "Group name (optional)")

	receiveCmd.MarkFlagRequired("server")
	receiveCmd.MarkFlagRequired("channel")
}
