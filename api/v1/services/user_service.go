package service

import (
	"context"
	"errors"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"

	"github.com/iacopoghilardi/mynance-service-api/internal/database"
	"github.com/iacopoghilardi/mynance-service-api/models"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) CreateUser(ctx context.Context, userData *models.User) error {
	db := database.GetDB().WithContext(ctx)
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    userData.Email,
		Password: hashedPassword,
	}
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	db := database.GetDB().WithContext(ctx)
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
	db := database.GetDB().WithContext(ctx)

	var userToUpdate models.User
	result := db.First(&userToUpdate, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	userToUpdate.Email = user.Email
	userToUpdate.Password = user.Password

	result = db.Save(&userToUpdate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	db := database.GetDB().WithContext(ctx)
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	db := database.GetDB().WithContext(ctx)
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
