package service

import (
	"github.com/hipeday/upay/internal/entities"
	"github.com/hipeday/upay/internal/repository"
	"time"
)

type TokenServiceImpl struct {
	TokenService
	repository repository.TokenRepository
}

func (a *TokenServiceImpl) Setup(repository repository.TokenRepository) {
	a.repository = repository
}

func (a *TokenServiceImpl) Save(token *entities.Token) error {
	queryToken, err := a.repository.SelectByTargetId(token.TargetId, token.Type)
	if err != nil {
		return err
	}
	if queryToken != nil {
		// 编辑token
		token.CreateAt = queryToken.CreateAt
		token.ID = queryToken.ID
		return a.repository.UpdateById(token)
	}
	return a.repository.Insert(token)
}

func (a *TokenServiceImpl) IsValidToken(accessToken string) (*entities.Token, bool, error) {
	tokenRepository := a.repository
	token, err := tokenRepository.SelectByAccessToken(accessToken)
	if err != nil {
		return nil, false, err
	}
	if token == nil {
		return nil, false, nil
	}

	if !token.ExpiresAt.After(time.Now()) {
		return nil, false, nil
	}
	return token, true, nil
}
