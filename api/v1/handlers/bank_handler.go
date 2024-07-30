package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
)

type BankHandler struct {
	Service *service.BankService
}

func NewBankHandler(s *service.BankService) *BankHandler {
	return &BankHandler{Service: s}
}

func (h *BankHandler) GetAllBanks(c *gin.Context) {
	fmt.Println("Banks list")

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(""))
}
