package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB    PostgresConfig
	OAuth GoogleOAuthConfig
	JWT   JWTConfig
}

type PostgresConfig struct {
	Url string
}

type GoogleOAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type JWTConfig struct {
	JWTSecret string
}


func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: PostgresConfig{
			Url: os.Getenv("POSTGRES_URL"),
		},	
		OAuth: GoogleOAuthConfig{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
		JWT: JWTConfig{	
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
	}

	return cfg, nil
}