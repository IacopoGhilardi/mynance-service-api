package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/internal/database"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"net/http"
)

var logger = utils.Logger

func main() {
	logger.Info("Init server")

	database.ConnectToDb()
	defer database.CloseDb()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	if err := r.Run(); err != nil {
		logger.Error("Server failed to start: ", err)
	} else {
		logger.Info("Server started succesfully")
	}
}
