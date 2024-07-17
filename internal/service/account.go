package service

import (
	"github.com/hipeday/upay/internal/constants"
	errors2 "github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/middleware"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
	"github.com/hipeday/upay/pkg/util"
	"time"
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

	var expiresAt = time.Minute * constants.TokenValidityPeriod

	// generate token
	token, err := util.GenerateToken(account.ID, account.Secret, &expiresAt)
	if err != nil {
		return nil, err
	}

	// 添加Token缓存
	middleware.TokenStoreCache.SetToken(token, account.Secret)

	return &response.SignIn{AccessToken: token}, nil
}
