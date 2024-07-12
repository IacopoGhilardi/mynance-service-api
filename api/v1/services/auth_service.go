package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	appConfig "github.com/iacopoghilardi/mynance-service-api/internal/config"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
	"github.com/iacopoghilardi/mynance-service-api/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	passwordMatch := utils.ComparePasswords(user.Password, password)
	if !passwordMatch {
		return "", errors.New("Wrong credentials")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) generateToken(user models.User) (string, error) {
	configs := appConfig.AppConfig

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JwtSecret))
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Email:    email,
		Password: string(hashedPassword),
		Profile:  models.Profile{},
	}

	if err := V1Services.UserService.CreateUser(context.Background(), &newUser); err != nil {
		return nil, err
	}

	newUser.CheckProfileComplete()
	return &newUser, nil
}

func verifyToken(tokenString string) error {
	configs := appConfig.AppConfig
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return configs.JwtSecret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}
