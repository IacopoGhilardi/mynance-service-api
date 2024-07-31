package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
)

func SetupBankRoutes(version *gin.RouterGroup) {
	bankRoutes := version.Group("/banks")
	bankHandler := handlers.V1Handlers.BankHandler

	bankRoutes.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	bankRoutes.GET("/", bankHandler.GetAllBanks)
	bankRoutes.POST("/", bankHandler.GetBankToken)
	// bankRoutes.POST("/", bankHandler.GetAllBanks)
	// bankRoutes.POST("/", bankHandler.GetAllBanks)

}
