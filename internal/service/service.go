package service

import "github.com/hipeday/upay/internal/repository"

type Service interface {
	Setup(repository repository.Repository)
}
