package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/MuxSphere/microkit/api-gateway/handlers"
	"github.com/MuxSphere/microkit/api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestEnv() {
	os.Setenv("PORT", "8080")
	os.Setenv("SERVICE_A_URL", "http://test-service-a:8080")
	os.Setenv("SERVICE_B_URL", "http://test-service-b:8080")
	os.Setenv("RATE_LIMIT", "10")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("CONSUL_ADDR", "test-consul:8500")
}

func TestConfigLoad(t *testing.T) {
	setupTestEnv()

	cfg, err := config.Load()
	assert.NoError(t, err)
	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "http://test-service-a:8080", cfg.ServiceAURL)
	assert.Equal(t, "http://test-service-b:8080", cfg.ServiceBURL)
	assert.Equal(t, 10, cfg.RateLimit)
	assert.Equal(t, "test-secret", cfg.JWTSecret)
	assert.Equal(t, "test-consul:8500", cfg.ConsulAddr)
}

func TestMiddlewareLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.Logger())

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMiddlewareRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.RateLimiter(2)) // A low limit for testing

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if i < 2 {
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			assert.Equal(t, http.StatusTooManyRequests, w.Code)
		}
	}
}

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	cfg := &config.Config{
		ServiceAURL: "http://test-service-a:8080",
		ServiceBURL: "http://test-service-b:8080",
	}

	handlers.SetupRoutes(r, cfg)

	// Test Service A route
	req, _ := http.NewRequest("GET", "/service-a/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test Service B route
	req, _ = http.NewRequest("GET", "/service-b/test", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test health check route
	req, _ = http.NewRequest("GET", "/health", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMainComponents(t *testing.T) {
	setupTestEnv()

	// Test config loading
	cfg, err := config.Load()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Test Gin setup
	gin.SetMode(gin.TestMode)
	r := gin.New()
	assert.NotNil(t, r)

	// Test middleware setup
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.RateLimiter(cfg.RateLimit))

	// Test route setup
	handlers.SetupRoutes(r, cfg)

	// Make a test request to ensure everything is set up correctly
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
