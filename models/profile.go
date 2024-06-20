package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID            uint      `json:"user_id" gorm:"uniqueIndex;not null"`
	Username          string    `json:"username" gorm:"type:varchar(50);uniqueIndex"`
	FirstName         string    `json:"first_name" gorm:"type:varchar(50)"`
	LastName          string    `json:"last_name" gorm:"type:varchar(50)"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	PhoneNumber       string    `json:"phone_number" gorm:"type:varchar(20)"`
	Address           string    `json:"address" gorm:"type:text"`
	City              string    `json:"city" gorm:"type:varchar(50)"`
	State             string    `json:"state" gorm:"type:varchar(50)"`
	PostalCode        string    `json:"postal_code" gorm:"type:varchar(20)"`
	Country           string    `json:"country" gorm:"type:varchar(50)"`
	ProfilePictureURL string    `json:"profile_picture_url" gorm:"type:text"`
}
