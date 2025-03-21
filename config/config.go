package config

import (
	// this will automatically load your .env file:

	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB    PostgresConfig
}

type PostgresConfig struct {
	Username string
	Password string
	Hostname string
	DB       string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			Hostname: os.Getenv("POSTGRES_HOST"),
			DB:       os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg, nil
}