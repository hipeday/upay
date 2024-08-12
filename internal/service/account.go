package service

import (
	errors2 "github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
	"github.com/hipeday/upay/pkg/util"
	token2 "github.com/hipeday/upay/pkg/util/token"
)

type AccountServiceImpl struct {
	repository repository.Repository
}

func (a *AccountServiceImpl) Setup(repository repository.Repository) {
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

	// save tokens

	return &response.SignIn{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
