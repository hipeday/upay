package service

import (
	"github.com/hipeday/upay/internal/entities"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
)

type Service[R repository.Repository] interface {
	Setup(repository R)
}

type AccountService interface {
	Service[repository.AccountRepository]
	SignIn(payload request.SignInPayload) (*response.SignIn, error)
}

type TokenService interface {
	Service[repository.TokenRepository]
	Save(token *entities.Token) error
	IsValidToken(token string) (*entities.Token, bool, error)
}

type SettingsService interface {
	Service[repository.SettingsRepository]
	Save(payload request.SaveSettingsPayload) (*response.SaveSettings, error)
	ModifyByConfigKey(payload request.SaveSettingsPayload) (*response.SaveSettings, error)
}
