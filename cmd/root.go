/*
Copyright © 2024 BUSINGE BISANGA <busingebisanga99@gmai.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type (
	SendService interface {
		SendMessage(topic, message string) error
	}

	ReceiveService interface {
		ReadMessages(topic string) error
	}

	// CobraCommander GI JOE!
	CobraCommander struct {
		rootCmd        *cobra.Command
		sendCfg        sendConfig
		sendService    SendService
		receiveCfg     receiveConfig
		receiveService ReceiveService
	}
)

func NewCobraCommander(sendService SendService, receiveService ReceiveService) *CobraCommander {
	return &CobraCommander{
		// rootCmd represents the base command when called without any subcommands
		rootCmd: &cobra.Command{
			Use:   "Kafka_and_CLIs",
			Short: "Command-line driven program that allows message exchange",
			Long: `This application allows you to send or receive messages from kafka cluster using the command line interface
		
					Cobra is a CLI library for Go that empowers applications.
					This application is a tool to generate the needed files
					to quickly create a Cobra application.`,
			// Uncomment the following line if your bare application
			// has an action associated with it:
			// Run: func(cmd *cobra.Command, args []string) { },
		},
		sendService:    sendService,
		receiveService: receiveService,
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (cbr *CobraCommander) Execute() {
	err := cbr.rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Kafka_and_CLIs.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

func SetupRootCmd(cbr *CobraCommander) {
	cbr.rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
