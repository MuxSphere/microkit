package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/MuxSphere/microkit/service-a/config"
	"github.com/MuxSphere/microkit/service-a/handlers"
	"github.com/MuxSphere/microkit/shared/database"
	"github.com/MuxSphere/microkit/shared/discovery"
	"github.com/MuxSphere/microkit/shared/logger"
	"github.com/MuxSphere/microkit/shared/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(httpRequestsTotal)
}

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
		l.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Initialize RabbitMQ
	rabbitMQ, err := rabbitmq.New(cfg.RabbitMQURL, l)
	if err != nil {
		l.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer rabbitMQ.Close()

	// 1st example: Publish a message
	err = rabbitMQ.PublishMessage("example_exchange", "example_routing_key", []byte("Hello, RabbitMQ!"))
	if err != nil {
		l.Error("Failed to publish message", zap.Error(err))
	}

	// 2nd example: Consume messages
	err = rabbitMQ.ConsumeMessages("example_queue", func(body []byte) error {
		l.Info("Received message", zap.ByteString("body", body))
		return nil
	})
	if err != nil {
		l.Error("Failed to consume messages", zap.Error(err))
	}

	// Sets up service discovery
	sd, err := discovery.NewServiceDiscovery(cfg.ConsulAddr)
	if err != nil {
		l.Fatal("Failed to create service discovery client", zap.Error(err))
	}

	// Register service with Consul
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		l.Fatal("Failed to convert port to int", zap.Error(err))
	}

	err = sd.RegisterService(cfg.ServiceName, cfg.Host, port)
	if err != nil {
		l.Fatal("Failed to register service with Consul", zap.Error(err))
	}
	defer sd.DeregisterService(cfg.ServiceName, cfg.Host, port) // Deregister on shutdown

	// Initialize Gin router and Prometheus middleware
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.GinMiddleware(l))
	r.Use(prometheusMiddleware())

	// Register Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Register routes
	handlers.RegisterRoutes(r, db, l)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Start HTTP server
	go func() {
		l.Info("Starting HTTP server", zap.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Failed to start HTTP server", zap.Error(err))
		}
	}()

	// Start gRPC server
	go func() {
		if err := startGRPCServer(l, cfg.GRPCPort); err != nil {
			l.Fatal("Failed to start gRPC server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", zap.Error(err))
	}

	l.Info("Server exiting")
}

// Prometheus middleware to track HTTP requests
func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
	}
}
