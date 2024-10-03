package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MuxSphere/microkit/service-a/config"
	"github.com/MuxSphere/microkit/service-a/handlers"
	"github.com/MuxSphere/microkit/shared/database"
	"github.com/MuxSphere/microkit/shared/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	l := logger.New(cfg.LogLevel)

	// Initialize database connection
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		l.Fatal("Failed to connect to database", "error", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.GinMiddleware(l))

	// Register routes
	handlers.RegisterRoutes(r, db, l)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Start server
	go func() {
		l.Info("Starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Failed to start server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", "error", err)
	}

	l.Info("Server exiting")
}
