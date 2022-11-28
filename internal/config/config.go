package config

import "os"

type Config struct {
	Name string
	Host string
	Port string
}

func NewConfig() *Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return &Config{
		Host: "127.0.0.1",
		Port: port,
	}
}
