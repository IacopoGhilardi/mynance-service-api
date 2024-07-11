package database

import (
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = utils.Logger
var db *gorm.DB

func ConnectToDb(connectionStr string) error {
	logger.Info("Connecting to db")
	var err error

	db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to open the DB connection: ", err)
		return err
	}

	logger.Info("Successfully connected to the database")
	return nil
}

func CloseDb() {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Error getting the underlying SQL DB: ", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.Error("Error closing the database: ", err)
	} else {
		logger.Info("Database connection closed")
	}
}

func GetDB() *gorm.DB {
	return db
}
