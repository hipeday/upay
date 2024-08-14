package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/constants/http"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/middleware"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/service"
	"github.com/jmoiron/sqlx"
)

type SettingsRouteImpl struct {
	service service.SettingsService
}

func setupSettings(db *sqlx.DB, engine *gin.Engine) {
	serviceInstance := service.GetSettingsServiceInstance(db)
	route := new(SettingsRouteImpl)
	route.Setup(serviceInstance)
	route.Register(engine, middleware.RequestLoggingMiddleware(), middleware.BearerAuthorizationMiddleware())
}

func (s *SettingsRouteImpl) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	group := engine.Group("", middlewares...)
	group.POST(addSettings, s.save)
	group.PUT(modifySettings, s.modify)
}

func (s *SettingsRouteImpl) Setup(service service.SettingsService) {
	s.service = service
}

func (s *SettingsRouteImpl) save(c *gin.Context) {
	var (
		payload request.SaveSettingsPayload
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	accountIdSource, exists := c.Get(http.AccountIdContext)
	if !exists {
		panic(errors.NewInternalServerError("The authentication is abnormal"))
	}

	payload.OperatorId = accountIdSource.(int64)
	response, err := s.service.Save(payload)
	if err != nil {
		panic(err)
	}

	c.JSON(200, response)
}

func (s *SettingsRouteImpl) modify(c *gin.Context) {
	var (
		payload request.SaveSettingsPayload
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	accountIdSource, exists := c.Get(http.AccountIdContext)
	if !exists {
		panic(errors.NewInternalServerError("The authentication is abnormal"))
	}

	payload.OperatorId = accountIdSource.(int64)
	response, err := s.service.ModifyByConfigKey(payload)
	if err != nil {
		panic(err)
	}

	c.JSON(200, response)
}
