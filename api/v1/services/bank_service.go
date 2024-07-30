package service

import "gorm.io/gorm"

type BankService struct {
	db *gorm.DB
}

func NewBankService(db *gorm.DB) *BankService {
	return &BankService{db: db}
}
