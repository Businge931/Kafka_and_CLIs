/*
Copyright © 2024 BUSINGE BISANGA <busingebisanga99@gmai.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Kafka_and_CLIs",
	Short: "Command-line driven program that allows message exchange",
	Long: `This application allows you to send or receive messages from kafka cluster using the command line interface

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
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

func SetupRootCmd() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
