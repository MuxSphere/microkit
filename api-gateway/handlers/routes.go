package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/MuxSphere/microkit/api-gateway/config"
	"github.com/MuxSphere/microkit/shared/discovery"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	// Service discovery with Consul
	sd, err := discovery.NewServiceDiscovery(cfg.ConsulAddr)
	if err != nil {
		panic(err)
	}

	// Health check route
	r.GET("/health", healthCheck)

	// Proxies for service-a and service-b
	r.Any("/service-a/*path", createServiceProxy(sd, "service-a"))
	r.Any("/service-b/*path", createServiceProxy(sd, "service-b"))
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createServiceProxy(sd *discovery.ServiceDiscovery, serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Discover the service through Consul
		service, err := sd.DiscoverService(serviceName)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
			return
		}

		// Creates the proxy for the discovered service
		targetURL := fmt.Sprintf("http://%s:%d", service.Address, service.Port)
		url, _ := url.Parse(targetURL)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
