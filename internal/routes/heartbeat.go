package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/routes/response"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type HeartbeatRoute struct {
}

func setupHeartbeat(db *sqlx.DB, engine *gin.Engine) {
	route := HeartbeatRoute{}
	route.Register(engine)
}

func (h HeartbeatRoute) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	engine.GET(ping, func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Ping{Message: "pong"})
	})
}

func (h HeartbeatRoute) Setup(service *service.Service) {

}
