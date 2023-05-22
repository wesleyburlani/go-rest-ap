package config

import (
	"log"
	"strings"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
)

const (
	DebugMode   string = "debug"
	ReleaseMode string = "release"
)

type Config struct {
	Mode     string `env:"MODE" envDefault:"debug"`
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

	config.Mode = strings.ToLower(config.Mode)

	if !slices.Contains([]string{DebugMode, ReleaseMode}, config.Mode) {
		config.Mode = DebugMode
	}

	return &config
}
