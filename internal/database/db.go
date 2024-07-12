package database

import (
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var log = utils.Logger
var db *gorm.DB

func ConnectToDb(connectionStr string) error {
	log.Info("Connecting to db")
	var err error

	db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Error("Failed to open the DB connection: ", err)
		return err
	}

	log.Info("Successfully connected to the database")
	return nil
}

func CloseDb() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("Error getting the underlying SQL DB: ", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Error("Error closing the database: ", err)
	} else {
		log.Info("Database connection closed")
	}
}

func GetDB() *gorm.DB {
	return db
}
