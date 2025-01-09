package apiserver

import "my_http/internal/app/store"

type Config struct{
	Port string `json:"port"`
	LogLevel string `json:"loglevel"`
	Store *store.Config `json:"store"`
}

func NewConfig() *Config{
	return &Config{
		Port: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}