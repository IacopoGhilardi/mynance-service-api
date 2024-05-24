package database

import (
	"database/sql"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
)

var logger = utils.Logger
var db *sql.DB

func ConnectToDb() {
	//Todo: prendere la string dalle configs
	logger.Info("Connecting to db")
	connStr := ""

	logger.Info("Connecting to db")

	if connStr == "" {
		logger.Info("Not connected")
		return
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to open the DB connection: ", err)
		return
	}

	if err := db.Ping(); err != nil {
		logger.Error("Failed to connect to the database: ", err)
		return
	}

	logger.Info("Successfully connected to the database")
}

func CloseDb() {
	if db != nil {
		err := db.Close()
		if err != nil {
			logger.Error("Error closing the database: ", err)
		} else {
			logger.Info("Database connection closed")
		}
	}
}
