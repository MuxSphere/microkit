package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func RegisterRoutes(r *gin.Engine, db *sqlx.DB, logger *zap.Logger) {
	r.GET("/health", healthCheck)

	// Add more routes here
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
