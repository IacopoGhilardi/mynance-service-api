package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/middleware"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
)

func SetupUserRoutes(version *gin.RouterGroup) {
	userRoutes := version.Group("/users")

	userHandler := handlers.V1Handlers.UserHandler
	authMiddleware := middleware.AuthMiddleware(service.V1Services.AuthService)

	userRoutes.Use(authMiddleware)
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
		userRoutes.GET("/", userHandler.GetAllUsers)

		profileHandler := handlers.V1Handlers.ProfileHandler

		userRoutes.GET("/:id/profile", profileHandler.GetProfile)
		userRoutes.POST("/:id/profile", profileHandler.UpdateProfile)
		userRoutes.PUT("/:id/profile", profileHandler.UpdateProfile)
	}
}
