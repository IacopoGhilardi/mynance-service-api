package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
)

func SetupUserRoutes(version *gin.RouterGroup) {
	userRoutes := version.Group("/users")

	userHandler := handlers.V1Handlers.UserHandler

	userRoutes.POST("/", userHandler.CreateUser)
	userRoutes.GET("/:id", userHandler.GetUser)
	userRoutes.PUT("/:id", userHandler.UpdateUser)
	userRoutes.DELETE("/:id", userHandler.DeleteUser)
	userRoutes.GET("/", userHandler.GetAllUsers)

	//Profile
	//userRoutes.GET("/:id/profile", handler.GetProfile)
	//userRoutes.GET("/:id/profile", handler.GetProfile)
}
