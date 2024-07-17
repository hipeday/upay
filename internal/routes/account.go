package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/logging"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type AccountRoute struct {
	service service.Service
}

func setupAccount(db *sqlx.DB, engine *gin.Engine) {
	accountRepository := repository.AccountRepositoryImpl{}
	accountRepository.Setup(db)

	accountService := service.AccountServiceImpl{}
	accountService.Setup(&accountRepository)

	route := AccountRoute{}
	route.Setup(&accountService)
	route.Register(engine)
}

func (a *AccountRoute) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	whitelist := engine.Group("", middlewares...)
	whitelist.POST(signin, a.signIn)
}

func (a *AccountRoute) Setup(service service.Service) {
	a.service = service
}

// signIn 登录
func (a *AccountRoute) signIn(c *gin.Context) {
	var (
		payload request.SignInPayload
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	logging.Logger().Debugf("SignIn payload: %+v", payload)

	signIn, err := a.service.SignIn(payload)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, signIn)
}
