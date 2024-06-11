package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
)

func SetupUserRoutes(version *gin.RouterGroup, handler *userHandler.UserHandler) {
	userRoutes := version.Group("/users")

	userRoutes.POST("/", handler.CreateUser)
	userRoutes.GET("/:id", handler.GetUser)
	userRoutes.PUT("/:id", handler.UpdateUser)
	userRoutes.DELETE("/:id", handler.DeleteUser)
	userRoutes.GET("/", handler.GetAllUsers)
}
