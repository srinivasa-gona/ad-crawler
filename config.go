package main

import (
	"ad-crawler/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func getDefaultConfig() (model.Configuration, error) {
	configuration := model.Configuration{}

	viper.SetDefault("ApplicationPort", "8080")
	viper.SetDefault("DBFileLocation", "ads_txt.db")
	viper.SetDefault("PublishersFileLocation", "publisher_data.csv")

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error in reading viper configuration %v", err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	err = viper.Unmarshal(&configuration)
	return configuration, err
}
