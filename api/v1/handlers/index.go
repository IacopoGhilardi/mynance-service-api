package handlers

import (
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
)

type V1ApiHandlers struct {
	UserHandler    *UserHandler
	ProfileHandler *ProfileHandler
	AuthHandler    *AuthHandler
}

var V1Handlers V1ApiHandlers

func InitHandlers() {
	V1Handlers = V1ApiHandlers{
		UserHandler:    NewUserHandler(service.V1Services.UserService),
		ProfileHandler: NewProfileHandler(service.V1Services.ProfileService),
		AuthHandler:    NewAuthHandler(service.V1Services.AuthService),
	}
}
