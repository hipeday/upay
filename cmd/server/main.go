package server

import (
	"github.com/hipeday/upay/config"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes"
	"log"
)

func Run() {
	// load and get config
	cfg := config.GetCfg()

	// init mysql db connection
	db, err := repository.InitMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init routes
	routes.SetupRouter(db, cfg)
}
