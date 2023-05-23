package config

import (
	"log"
	"os"
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
	ServiceName string `env:"SERVICE_NAME" envDefault:"service-name"`
	Mode        string `env:"MODE" envDefault:"debug"`
	HttpHost    string `env:"HTTP_HOST" envDefault:"localhost"`
	HttpPort    int    `env:"HTTP_PORT" envDefault:"8080"`
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
	config := Config{}

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	config.Mode = strings.ToLower(config.Mode)

	possibleModes := []string{DebugMode, ReleaseMode}
	if !slices.Contains(possibleModes, config.Mode) {
		log.Fatalf("config MODE must be one of: [%s]", strings.Join(possibleModes, ","))
	}

	return &config
}
