package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/service"
)

type Route[R repository.Repository, S service.Service[R]] interface {
	Register(engine *gin.Engine, middlewares ...gin.HandlerFunc)
	Setup(service S)
}

type IAccountRoute interface {
	Route[repository.AccountRepository, service.AccountService]
	signIn(c *gin.Context)
}

type AccountRoute struct {
	service service.AccountService
}
