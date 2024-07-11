package main

import (
	"fmt"
	"github.com/iacopoghilardi/mynance-service-api/internal/app"
	"github.com/iacopoghilardi/mynance-service-api/internal/config"
	appConfig "github.com/iacopoghilardi/mynance-service-api/internal/config"
	"github.com/iacopoghilardi/mynance-service-api/internal/database"
)

func main() {

	var err error
	err = config.InitConfig()
	if err != nil {
		panic(err)
	}

	configs := appConfig.AppConfig
	host := configs.DBHost
	port := configs.DBPort
	user := configs.DBUser
	password := configs.DBPass
	dbname := configs.DBName

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Europe/Rome", host, user, password, dbname, port)

	err = database.ConnectToDb(connStr)
	if err != nil {
		panic(err)
	}
	defer database.CloseDb()
	err = app.InitApp()
	if err != nil {
		panic(err)
	}
}
