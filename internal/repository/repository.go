package repository

import (
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AccountRepository
	Setup(db *sqlx.DB)
	TableName() string
	Columns() []string
	Columns2Query() string
}

type AccountRepository interface {
	SelectAccountByUsername(username string) (*entities.Account, error)
	SelectSignIn(username string, password string) (*entities.Account, error)
}
