package app

import (
	v1Handlers "github.com/iacopoghilardi/mynance-service-api/api/v1/handlers"
	v1Services "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/config"
	"github.com/iacopoghilardi/mynance-service-api/internal/database"
	"github.com/iacopoghilardi/mynance-service-api/internal/router"
	"github.com/iacopoghilardi/mynance-service-api/pkg/utils"
)

var logger = utils.Logger

func InitApp() error {
	var err error
	logger.Info("Init server")

	err = config.InitConfig()
	if err != nil {
		panic(err)
	}

	err = database.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer database.CloseDb()

	v1Services.InitServices()
	v1Handlers.InitHandlers()

	err = router.InitRouter()
	if err != nil {
		panic(err)
	}

	return err

}
