package service

import (
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
)

type Service interface {
	AccountService
	Setup(repository repository.Repository)
}

type AccountService interface {
	SignIn(payload request.SignInPayload) (*response.SignIn, error)
}
