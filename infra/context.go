package infra

import (
	"car-rental/config"
	"car-rental/controllers"
	"car-rental/databases"
	"car-rental/log"
	"car-rental/repositories"
	"car-rental/services"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Log *zap.Logger
	Ctl controllers.Controller
}

func NewServiceContext() (*ServiceContext, error) {
	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	appConfig, err := config.NewConfig()
	if err != nil {
		return handlerError(err)
	}

	logger, err := log.InitLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	db, err := databases.InitDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	repo := repositories.NewRepository(db, logger)
	service := services.NewService(repo, logger)
	ctl := controllers.NewController(*service, logger)

	return &ServiceContext{
		Cfg: appConfig,
		Log: logger,
		Ctl: *ctl,
	}, nil
}
