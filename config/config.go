package config

import (
	"log"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpHost string `env:"HTTP_HOST" envDefault:"localhost"`
	HttpPort int    `env:"HTTP_PORT" envDefault:"8080"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	config := Config{}

	err = env.Parse(&config)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return &config
}
