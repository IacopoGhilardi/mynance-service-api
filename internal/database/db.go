package database

import (
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = utils.Logger
var db *gorm.DB

func ConnectToDb() error {
	//Todo: prendere la string dalle configs
	logger.Info("Connecting to db")
	connStr := ""
	//host := config.AppConfig.DBHost
	//port := config.AppConfig.DBPort
	//user := config.AppConfig.DBUser
	//password := config.AppConfig.DBPass
	//dbname := config.AppConfig.DBName

	//connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	//	host, user, password, dbname, port)

	var err error
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable TimeZone=Europe/Rome"

	if connStr == "" {
		logger.Info("Not connected")
		return nil
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
