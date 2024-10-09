package main

import (
	"log"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/MuxSphere/microkit/api-gateway/handlers"
	"github.com/MuxSphere/microkit/api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a production logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Set up Gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))

	// Add rate limiting middleware
	r.Use(middleware.RateLimiter(cfg.RateLimit))

	// Set up routes
	handlers.SetupRoutes(r, cfg)

	// Start server
	logger.Info("Starting API Gateway", zap.String("port", cfg.Port))
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
