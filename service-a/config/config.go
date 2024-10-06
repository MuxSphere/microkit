package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	GRPCPort    string
	DatabaseURL string
	RabbitMQURL string
	LogLevel    string
	ServiceName string
	Host        string
	ConsulAddr  string
}

func Load() (*Config, error) {
	// Loads .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Automatically load environment variables
	viper.AutomaticEnv()

	// Set default values in case env vars are not set
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("DATABASE_URL", "postgres://user:password@localhost:5432/microservices?sslmode=disable")
	viper.SetDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("SERVICE_NAME", "service-a")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("CONSUL_ADDR", "consul:8500")

	// Creates and populate Config struct from environment variables
	var cfg Config
	cfg.Port = viper.GetString("PORT")
	cfg.GRPCPort = viper.GetString("GRPC_PORT")
	cfg.DatabaseURL = viper.GetString("DATABASE_URL")
	cfg.RabbitMQURL = viper.GetString("RABBITMQ_URL")
	cfg.LogLevel = viper.GetString("LOG_LEVEL")
	cfg.ServiceName = viper.GetString("SERVICE_NAME")
	cfg.Host = viper.GetString("HOST")
	cfg.ConsulAddr = viper.GetString("CONSUL_ADDR")

	return &cfg, nil
}
