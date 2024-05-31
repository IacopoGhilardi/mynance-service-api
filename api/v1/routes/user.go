package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/iacopoghilardi/mynance-service-api/api/v1/handlers/user"
)

func SetupUserRoutes(version *gin.RouterGroup) {
	userRoutes := version.Group("/users")

	userRoutes.POST("/", userHandler.GetUser)
	userRoutes.GET("/", userHandler.GetUser)
}
