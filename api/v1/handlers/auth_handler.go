package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers/dto"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) TestEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(""))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse("Invalid request payload"))
		return
	}

	fmt.Printf("Ecco la password: %v", req.Password)
	fmt.Printf(" Ecco la email: %v", req.Email)

	token, err := h.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse("Invalid credentials"))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(token))
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse("Invalid request payload"))
		return
	}

	user, err := h.Service.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GenerateErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(user))
}
