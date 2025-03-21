package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB    PostgresConfig
}

type PostgresConfig struct {
	Url string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: PostgresConfig{
			Url: os.Getenv("POSTGRES_URL"),
		},
	}

	return cfg, nil
}