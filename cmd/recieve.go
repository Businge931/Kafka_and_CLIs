package cmd

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type receiveConfig struct {
	receiveServer string
	receiveTopic  string
	startFrom     string
	receiveGroup  string
	dynamicGroup  bool
}

// SetupReceiveCmd sets up the flags and adds the command to the root
func (cbr *CobraCommander) SetupReceiveCmd() {
	// receiveCmd defines the Cobra command for consuming messages
	receiveCmd := &cobra.Command{
		Use:   "receive",
		Short: "Receive messages from Kafka",
		Run: func(_ *cobra.Command, _ []string) {
			// Echo the parameters to the user
			log.Printf("You are receiving from the channel: '%s'\n", cbr.receiveCfg.receiveTopic)
			log.Printf("You are receiving from the '%s'\n", cbr.receiveCfg.startFrom)
			log.Printf("You are receiving through the server: '%s'\n", cbr.receiveCfg.receiveServer)

			if cbr.receiveCfg.receiveGroup != "" {
				log.Printf("You are part of the receiving group: '%s'\n", cbr.receiveCfg.receiveGroup)
			}

			// Check for dynamic group flag
			if cbr.receiveCfg.dynamicGroup {
				// Create a unique consumer group name if dynamicGroup is true
				cbr.receiveCfg.receiveGroup = fmt.Sprintf("%s-%d", cbr.receiveCfg.receiveGroup, time.Now().UnixNano())
				log.Printf("Dynamic group created: '%s'\n", cbr.receiveCfg.receiveGroup)
			}

			// Use the injected service to receive messages
			if err := cbr.receiveService.ReadMessages(cbr.receiveCfg.receiveTopic); err != nil {
				log.Fatalf("Failed to receive messages: %s", err)
			}
		},
	}

	cbr.rootCmd.AddCommand(receiveCmd)

	// Define flags for the receive command
	receiveCmd.Flags().StringVar(&cbr.receiveCfg.receiveServer, "server", "", "Kafka connection string (required)")
	receiveCmd.Flags().StringVar(&cbr.receiveCfg.receiveTopic, "channel", "", "Kafka topic (required)")
	receiveCmd.Flags().StringVar(&cbr.receiveCfg.startFrom, "from", "earliest", "Start consuming from (start|latest)")
	receiveCmd.Flags().StringVar(&cbr.receiveCfg.receiveGroup, "group", "", "Group name (optional)")
	receiveCmd.Flags().BoolVar(&cbr.receiveCfg.dynamicGroup, "dynamic-group", false, "Use dynamic group (optional)")

	err := receiveCmd.MarkFlagRequired("server")
	if err != nil {
		log.Printf("failed to define flag: %s", cbr.receiveCfg.receiveServer)
	}

	err = receiveCmd.MarkFlagRequired("channel")
	if err != nil {
		log.Printf("failed to define flag: %s", cbr.receiveCfg.receiveTopic)
	}
}
