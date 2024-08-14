package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/middleware"
	"github.com/hipeday/upay/pkg/config"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

func SetupRouter(db *sqlx.DB, cfg config.Config) {
	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()
	engine.Use(middleware.ErrorMiddleware())
	setup(db, engine)
	err := engine.Run(cfg.Server.IP + ":" + strconv.Itoa(int(cfg.Server.Port)))
	if err != nil {
		log.Fatal(err)
	}
}

func setup(db *sqlx.DB, engine *gin.Engine) {
	setupHeartbeat(db, engine)
	setupAccount(db, engine)
	setupSettings(db, engine)
	setupOrder(db, engine)
}
