package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string  `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string  `json:"-" gorm:"type:varchar(255);not null"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;"`
}
