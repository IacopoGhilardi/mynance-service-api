package service

import (
	"context"

	"github.com/iacopoghilardi/mynance-service-api/models"
	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID int64) (*models.Profile, error) {
	db := s.db.WithContext(ctx)
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, userID int64, updatedProfile *models.Profile) (*models.Profile, error) {
	db := s.db.WithContext(ctx)
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&profile).Updates(updatedProfile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
