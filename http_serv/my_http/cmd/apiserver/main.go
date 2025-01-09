package main

import (
	"flag"
	"log"
	"my_http/internal/app/apiserver"

	"github.com/spf13/viper"
)

var PathToConfig string

func init(){
	flag.StringVar(&PathToConfig, "config-path", "./configs/apiserver.json", "path to config file")
}

func main(){
	flag.Parse()

	config := apiserver.NewConfig()

	viper.SetConfigName("apiserver")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil{
		log.Fatal(err)
	}

	config.LogLevel = viper.GetString("loglevel")
	config.Port = viper.GetString("port")
	config.Store.DatabaseURL = viper.GetString("store.database_url")
	
	s := apiserver.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}