package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadAddConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.simple")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
