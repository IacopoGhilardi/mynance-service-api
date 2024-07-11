package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
)

func SetupAuthRoutes(version *gin.RouterGroup) {

	authHandler := handlers.V1Handlers.AuthHandler

	version.POST("/register", authHandler.TestEndpoint)
	version.POST("/login", authHandler.TestEndpoint)
	version.POST("/logout", authHandler.TestEndpoint)
	version.POST("/reset-password", authHandler.TestEndpoint)
}
