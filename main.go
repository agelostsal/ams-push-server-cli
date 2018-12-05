package main

import (
	"github.com/ARGOeu/ams-push-server-cli/cmd"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// initialize viper
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	// initialize the command chain and execute it
	cmd.NewRootCommand().Execute()
}
