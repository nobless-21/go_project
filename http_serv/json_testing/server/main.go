package main

import (
	"fmt"
	"log"
	"test_json/store"

	"github.com/spf13/viper"
)

type Config struct{
	LogLevel string `json:"loglevel"`
	Port string `json:"port"`
	Store *store.Config `json:"store"`
}

func NewConfig() *Config{
	return&Config{
		LogLevel: "debug",
		Port: "8080",
		Store: store.NewConfig(),
	}
}

func main(){
	json, err := loadConfig("configs")

	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println(json.LogLevel, json.Port, json.Store.DatabaseURL)
}

func loadConfig(filename string) (*Config, error){
	viper.SetConfigName(filename)
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil{
		return nil, err
	}

	var cfg Config

	cfg.Store = store.NewConfig()

	cfg.LogLevel = viper.GetString("loglevel")
	cfg.Port = viper.GetString("port")
	cfg.Store.DatabaseURL = viper.GetString("store.database_url")

	return &cfg, nil
}