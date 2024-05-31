package user

import (
	"context"
	"errors"

	"github.com/iacopoghilardi/mynance-service-api/internal/database"
	"github.com/iacopoghilardi/mynance-service-api/models"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user *models.User) error {
	db := database.GetDB().WithContext(ctx)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUser(ctx context.Context, id int64) (*models.User, error) {
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

func UpdateUser(ctx context.Context, user *models.User) error {
	db := database.GetDB().WithContext(ctx)
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(ctx context.Context, id int64) error {
	db := database.GetDB().WithContext(ctx)
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllUsers(ctx context.Context) ([]models.User, error) {
	db := database.GetDB().WithContext(ctx)
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
