package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/iacopoghilardi/mynance-service-api/api/v1/handlers/user"
)

func SetupAuthRoutes(versionGroup *gin.RouterGroup) {
	versionGroup.POST("/", userHandler.GetUser)
	versionGroup.GET("/", userHandler.GetUser)
}
