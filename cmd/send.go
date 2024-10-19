package cmd

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type sendConfig struct {
	sendKafkaServer string
	sendTopic       string
	sendGroup       string
}

func (cbr *CobraCommander) SetupSendCmd() {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send messages to Kafka",
		Run: func(_ *cobra.Command, _ []string) {
			// Echo the parameters to the user
			log.Printf("You have decided to send to the channel: '%s'\n", cbr.sendCfg.sendTopic)
			log.Printf("You are sending through the server: '%s'\n", cbr.sendCfg.sendKafkaServer)

			if cbr.sendCfg.sendGroup != "" {
				log.Printf("You are sending through the group: '%s'\n", cbr.sendCfg.sendGroup)
			}

			cbr.SendMessage(cbr.sendCfg.sendTopic)
		},
	}

	cbr.rootCmd.AddCommand(sendCmd)

	// Define flags for the send command
	sendCmd.Flags().StringVar(&cbr.sendCfg.sendKafkaServer, "server", "", "Kafka connection string (required)")
	sendCmd.Flags().StringVar(&cbr.sendCfg.sendTopic, "channel", "", "Kafka topic (required)")
	sendCmd.Flags().StringVar(&cbr.sendCfg.sendGroup, "group", "", "Group name (optional)")

	err := sendCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", cbr.sendCfg.sendKafkaServer)
	}

	err = sendCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", cbr.sendCfg.sendTopic)
	}
}

// SendMessage sends messages using the provided Service.
func (cbr *CobraCommander) SendMessage(topic string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("Enter message (or 'q' to quit): ")
		message, _ := reader.ReadString('\n')

		// Remove the newline character
		message = strings.TrimSpace(message)
		if message == "q" {
			break
		}

		// Send the message using the service
		if err := cbr.sendService.SendMessage(topic, message); err != nil {
			log.Errorf("Failed to send message: %v", err)
		}
	}
}
