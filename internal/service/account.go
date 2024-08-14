package service

import (
	"github.com/hipeday/upay/internal/entities"
	errors2 "github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
	"github.com/hipeday/upay/pkg/util"
	token2 "github.com/hipeday/upay/pkg/util/token"
	"time"
)

type AccountServiceImpl struct {
	AccountService
	repository repository.AccountRepository
}

func (a *AccountServiceImpl) Setup(repository repository.AccountRepository) {
	a.repository = repository
}

func (a *AccountServiceImpl) SignIn(payload request.SignInPayload) (*response.SignIn, error) {
	var (
		err               error
		accountRepository = a.repository
	)

	account, err := accountRepository.SelectAccountByUsername(payload.Username)

	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors2.NewIllegalArgumentError("username or password error")
	}

	password := util.MD5SaltHash(payload.Password, account.Secret)

	account, err = accountRepository.SelectSignIn(payload.Username, password)

	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors2.NewIllegalArgumentError("username or password error")
	}

	// generate access token
	accessToken, expiresAt, err := token2.GenerateAccessToken()
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, err := token2.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// save tokens If the data already exists for the current account, it should be overwritten
	tokenService := GetTokenServiceInstance(accountRepository.GetDB())

	now := time.Now()
	token := entities.Token{
		Entity: entities.Entity{
			ID:       0,
			CreateAt: &now,
		},
		TargetId:     account.ID,
		Type:         entities.AccountTokenType,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    &expiresAt,
	}
	err = tokenService.Save(&token)
	if err != nil {
		return nil, err
	}

	return &response.SignIn{AccessToken: accessToken, RefreshToken: refreshToken, ExpiresIn: expiresAt.UnixMilli() - time.Now().UnixMilli()}, nil
}
