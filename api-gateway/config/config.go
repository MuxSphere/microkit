package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	ServiceAURL string
	ServiceBURL string
	RateLimit   int
	JWTSecret   string
}

func Load() (*Config, error) {
	// Loads environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	viper.AutomaticEnv()

	viper.SetDefault("PORT", "8000")
	viper.SetDefault("SERVICE_A_URL", "http://service-a:8080")
	viper.SetDefault("SERVICE_B_URL", "http://service-b:8080")
	viper.SetDefault("RATE_LIMIT", 100)
	viper.SetDefault("JWT_SECRET", "your-secret-key")

	var cfg Config
	cfg.Port = viper.GetString("PORT")
	cfg.ServiceAURL = viper.GetString("SERVICE_A_URL")
	cfg.ServiceBURL = viper.GetString("SERVICE_B_URL")
	cfg.RateLimit = viper.GetInt("RATE_LIMIT")
	cfg.JWTSecret = viper.GetString("JWT_SECRET")

	return &cfg, nil
}
