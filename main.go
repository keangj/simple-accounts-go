package main

import (
	"log"
	"simple-accounts/cmd"

	"github.com/spf13/viper"
)

// @title           Simple Accounts API
// @version         1.0
// @description     This is a sample server for Simple Accounts API.

// @contact.name   Jay
// @contact.url    null
// @contact.email  keangjay@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Run()
}
