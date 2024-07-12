package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mynance-service-api/api/v1/handlers/dto"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
	"github.com/iacopoghilardi/mynance-service-api/models"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	if err := h.Service.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.GenerateSuccessResponse(user))
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}
	user, err := h.Service.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, utils.GenerateErrorResponse("User not found"))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(user))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var updateUserDto dto.UpdateUserRequestDto
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&updateUserDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	user := models.User{
		Email: updateUserDto.Email,
	}

	if err := h.Service.UpdateUser(c.Request.Context(), id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(updateUserDto))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var users []models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}
	if err := h.Service.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(users))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(users))
}
