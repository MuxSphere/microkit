package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("/health", healthCheck)

	serviceAProxy := createProxy(cfg.ServiceAURL)
	r.Any("/service-a/*path", serviceAProxy)

	serviceBProxy := createProxy(cfg.ServiceBURL)
	r.Any("/service-b/*path", serviceBProxy)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createProxy(targetURL string) gin.HandlerFunc {
	url, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
