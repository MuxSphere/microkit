package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
	GRPCPort    string
}

func Load() (*Config, error) {
	// Loads .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Set default values in case env vars are not set
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable")
	viper.SetDefault("LOG_LEVEL", "info")

	viper.AutomaticEnv()

	var cfg Config
	cfg.Port = viper.GetString("PORT")
	cfg.DatabaseURL = viper.GetString("DATABASE_URL")
	cfg.LogLevel = viper.GetString("LOG_LEVEL")

	return &cfg, nil
}
