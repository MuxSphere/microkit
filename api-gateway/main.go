package main

import (
	"log"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/MuxSphere/microkit/api-gateway/handlers"
	"github.com/MuxSphere/microkit/api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up Gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	// Add rate limiting middleware
	r.Use(middleware.RateLimiter(cfg.RateLimit))

	// Set up routes
	handlers.SetupRoutes(r, cfg)

	// Start server
	log.Printf("Starting API Gateway on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
