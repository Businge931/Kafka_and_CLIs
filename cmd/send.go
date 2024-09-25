package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Businge931/Kafka_and_CLIs/producer"
)

var (
	sendKafkaServer string
	sendTopic       string
	sendGroup       string

	// Injected Producer function for testing:  returns the interface producer.MessageProducer, not the concrete *KafkaProducer
	// CreateProducerFunc = func(server string) (producer.MessageProducer, error) {
	// 	return producer.NewKafkaProducer(server)
	// }
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

		// Create a new producer 
		// kp, err := CreateProducerFunc(sendKafkaServer)
		kp, err := producer.NewKafkaProducer(sendKafkaServer)
		if err != nil {
			log.Fatalf("Failed to create Kafka producer: %s", err)
		}
		defer kp.Close()

		// Call the ProduceMessages function
		producer.ProduceMessages(kp, sendTopic, false)
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
