package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/MuxSphere/microkit/api-gateway/handlers"
	"github.com/MuxSphere/microkit/api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func setupRouter() (*gin.Engine, *observer.ObservedLogs) {
	gin.SetMode(gin.TestMode)

	// Create a logger with an observer for testing
	core, logs := observer.New(zap.InfoLevel)
	logger := zap.New(core)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))

	cfg := &config.Config{
		RateLimit: 10, // Set a rate limit for testing
		Port:      "8080",
	}

	r.Use(middleware.RateLimiter(cfg.RateLimit))

	handlers.SetupRoutes(r, cfg)

	return r, logs
}

func TestHealthCheck(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestLogger(t *testing.T) {
	router, logs := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 1, logs.Len())
	logEntry := logs.All()[0]

	assert.Equal(t, "Request", logEntry.Message)
	assert.Equal(t, int64(http.StatusOK), logEntry.ContextMap()["status"])
	assert.Equal(t, "GET", logEntry.ContextMap()["method"])
	assert.Equal(t, "/health", logEntry.ContextMap()["path"])
	assert.Contains(t, logEntry.ContextMap(), "latency")
}

func TestRateLimiter(t *testing.T) {
	router, _ := setupRouter()

	for i := 0; i < 11; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)

		if i < 10 {
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			assert.Equal(t, http.StatusTooManyRequests, w.Code)
		}
	}

	// Wait for rate limiter to reset
	time.Sleep(time.Second)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInvalidRoute(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/invalid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
