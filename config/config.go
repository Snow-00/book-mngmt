package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT 		int
	DB_PORT 	int
	DB_NAME 	string
	DB_PASSWORD string
	DB_USER 	string
	DB_HOST 	string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Load config success")
}