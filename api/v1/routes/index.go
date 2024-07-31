package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	prefix := r.Group("/api/v1")

	prefix.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	SetupUserRoutes(prefix)
	SetupAuthRoutes(prefix)
	SetupBankRoutes(prefix)
}
