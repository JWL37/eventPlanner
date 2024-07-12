package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port         string `env:"REST_PORT" default:"8082"`
	DataBasePath string `env:"DATA_BASE" default:"postgres://postgres:postgres@postgres:5432/postgres"`
}

func ConfigInit(cfg *Config) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
}
