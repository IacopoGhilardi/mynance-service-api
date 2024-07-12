package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
)

func SetupAuthRoutes(version *gin.RouterGroup) {

	authHandler := handlers.V1Handlers.AuthHandler

	version.POST("/register", authHandler.Register)
	version.POST("/login", authHandler.Login)
	version.POST("/reset-password", authHandler.TestEndpoint)
}
