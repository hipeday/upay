package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/middleware"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type AccountRouteImpl struct {
	AccountRoute
	service service.AccountService
}

func (a *AccountRouteImpl) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	whitelist := engine.Group("", middlewares...)
	whitelist.POST(signin, a.signIn)
}

func (a *AccountRouteImpl) Setup(service service.AccountService) {
	a.service = service
}

func setupAccount(db *sqlx.DB, engine *gin.Engine) {
	serviceInstance := service.GetAccountServiceInstance(db)
	var route AccountRoute = new(AccountRouteImpl)
	route.Setup(serviceInstance)
	route.Register(engine, middleware.RequestLoggingMiddleware())
}

// signIn 登录
func (a *AccountRouteImpl) signIn(c *gin.Context) {
	var (
		payload request.SignInPayload
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	signIn, err := a.service.SignIn(payload)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, signIn)
}
