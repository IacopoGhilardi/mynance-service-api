package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"net/http"
)

var logger = utils.Logger

func GetUser(c *gin.Context) {
	fmt.Println("Get user")

	logger.Info("Getting user")

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
