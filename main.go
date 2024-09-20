/*
Copyright © 2024 BUSINGE BISANGA <busingebisanga99@gmail.com>
*/
package main

import (
	"github.com/Businge931/Kafka_and_CLIs/cmd"
)

func main() {
	// Call the setup functions to initialize commands
	cmd.SetupRootCmd()
	cmd.SetupReceiveCmd()
	cmd.SetupSendCmd()

	// Execute the root command
	cmd.Execute()
}
