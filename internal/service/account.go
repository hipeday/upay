package service

import "github.com/hipeday/upay/internal/repository"

type AccountService struct {
	repository repository.Repository
}

func (a AccountService) Setup(repository repository.Repository) {
	a.repository = repository
}
