package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
)

type AccountRoute struct {
	service service.Service
}

func setupAccount(db *sqlx.DB, engine *gin.Engine) {
	accountRepository := repository.AccountRepository{}
	accountRepository.Setup(db)

	accountService := service.AccountService{}
	accountService.Setup(accountRepository)

	route := AccountRoute{}
	route.Setup(accountService)
	route.Register(engine)
}

func (a AccountRoute) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	engine.GET("/quickstart", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello world"})
	})
}

func (a AccountRoute) Setup(service service.Service) {
	a.service = service
}
