package routes

import (
	"github.com/gin-gonic/gin"
	v1Handler "github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
)

func SetupRoutes(r *gin.Engine) {
	prefix := r.Group("/api/v1")

	prefix.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	SetupUserRoutes(prefix, v1Handler.NewUserHandler(&service.UserService{}))
}
