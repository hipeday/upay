package server

import (
	"github.com/hipeday/upay/internal/application"
	"github.com/hipeday/upay/internal/logging"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes"
	"github.com/hipeday/upay/pkg/config"
)

func Run() {
	// load and get config
	cfg := config.GetCfg()

	// setup logger
	logging.Setup(cfg.Logger)

	// setup application
	application.Setup(cfg.Application)

	// init mysql db connection
	db, err := repository.InitMySQL(cfg)
	if err != nil {
		logging.Logger().Fatal(err)
	}

	// init routes
	routes.SetupRouter(db, cfg)
}
