package app

import (
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

	err = router.InitRouter()
	if err != nil {
		panic(err)
	}

	return err

}
