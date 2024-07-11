package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/iacopoghilardi/mynance-service-api/api/v1/routes"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"os"
)

var logger = utils.Logger
var router *gin.Engine

func InitRouter() error {
	var err error
	router = gin.Default()

	v1.SetupRoutes(router)

	if os.Getenv("GIN_MODE") != "test" {
		if err := router.Run(); err != nil {
			logger.Error("Server failed to start: ", err)
			return err
		}
	}

	return err
}

func GetRouter() *gin.Engine {
	return router
}
