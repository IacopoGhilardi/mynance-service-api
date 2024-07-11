package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/iacopoghilardi/mynance-service-api/models"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	utils.Logger.Info("Creating new user: " + user.Email)
	db := s.db.WithContext(ctx)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	utils.Logger.Info("Created!")
	return nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	db := s.db.WithContext(ctx)
	var user models.User
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	db := s.db.WithContext(ctx)
	var existingUser models.User
	if err := db.First(&existingUser, id).Error; err != nil {
		return err
	}

	fmt.Println("Arrivi qua a fare l'update")
	existingUser.Email = user.Email

	fmt.Println("Arrivi qua a fare l'update: " + user.Email)
	existingUser.Password = user.Password

	if err := db.Save(&existingUser).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	db := s.db.WithContext(ctx)
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	db := s.db.WithContext(ctx)
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
