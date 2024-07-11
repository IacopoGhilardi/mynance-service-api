package service

import "github.com/iacopoghilardi/mynance-service-api/internal/database"

type V1ApiServices struct {
	UserService    *UserService
	ProfileService *ProfileService
	AuthService    *AuthService
}

var V1Services V1ApiServices

func InitServices() {
	db := database.GetDB()

	V1Services = V1ApiServices{
		UserService:    NewUserService(db),
		ProfileService: NewProfileService(db),
		AuthService:    NewAuthService(db),
	}

}
