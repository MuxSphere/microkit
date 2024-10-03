package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
}

func Load() (*Config, error) {
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
