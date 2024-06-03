package database

import (
	"fmt"

	appConfig "github.com/iacopoghilardi/mynance-service-api/internal/config"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = utils.Logger
var db *gorm.DB

func ConnectToDb() error {
	logger.Info("Connecting to db")
	configs := appConfig.AppConfig
	host := configs.DBHost
	port := configs.DBPort
	user := configs.DBUser
	password := configs.DBPass
	dbname := configs.DBName

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome", host, user, password, dbname, port)

	var err error
	if connStr == "" {
		logger.Info("Not connected")
		return nil
	}

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
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
