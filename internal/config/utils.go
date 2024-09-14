package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	versionKey = "version"
	currentListKey = "current_list"
	listIDKey = "list_id"
	taskIDKey = "task_id"
) 

func LoadConfig() {
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault(versionKey, "0.0.1")
	viper.SetDefault(currentListKey, 0)
	viper.SetDefault(listIDKey, 1)
	viper.SetDefault(taskIDKey, 1)
	if err := viper.SafeWriteConfigAs("config.yml"); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			log.Fatalf("Error while setup configuration: %s", err)
		}
	}
	viper.ReadInConfig()
	viper.WatchConfig()
}

func GetCurrentList() int {
	return viper.GetInt(currentListKey)
}

func GetListId() int {
	return viper.GetInt(listIDKey)
}

func CheckoutList(id int) {
	viper.Set(currentListKey, id)
	if err := viper.WriteConfig(); err != nil {	
		log.Fatalf("Cannot change current_list")
	}
}

func IncrementListId() {
	viper.Set(listIDKey, viper.GetInt(listIDKey) + 1)
	if err := viper.WriteConfig(); err != nil {	
		log.Fatalf("Cannot increment list_id")
	}
}