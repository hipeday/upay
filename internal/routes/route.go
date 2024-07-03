package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/config"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type Route interface {
	Register(engine *gin.Engine, middlewares ...gin.HandlerFunc)
	Setup(service *service.Service)
}

func SetupRouter(db *sqlx.DB, cfg config.Config) {
	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()
	setup(db, engine)
	err := engine.Run(cfg.Server.IP + ":" + strconv.Itoa(int(cfg.Server.Port)))
	if err != nil {
		log.Fatal(err)
	}
}

func setup(db *sqlx.DB, engine *gin.Engine) {
	setupHeartbeat(db, engine)
	setupAccount(db, engine)
	setupOrder(db, engine)
}
