package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
	"github.com/iacopoghilardi/mynance-service-api/models"
)

type ProfileHandler struct {
	Service *service.ProfileService
}

func NewProfileHandler(s *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{Service: s}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	profile, err := h.Service.GetProfileByUserID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, utils.GenerateErrorResponse("Profile not found"))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(profile))
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var profile models.Profile
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(err.Error()))
		return
	}

	updatedProfile, err := h.Service.UpdateProfile(c.Request.Context(), userID, &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateSuccessResponse(updatedProfile))
}
