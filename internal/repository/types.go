package repository

import (
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Setup(db *sqlx.DB)
	TableName() string
	Columns() []string
	Columns2Query() string
	GetDB() *sqlx.DB
}

type AccountRepository interface {
	Repository
	SelectAccountByUsername(username string) (*entities.Account, error)
	SelectSignIn(username string, password string) (*entities.Account, error)
}

type TokenRepository interface {
	Repository
	Insert(*entities.Token) error
	UpdateById(*entities.Token) error
	SelectByTargetId(targetId int64, tokenType entities.TokenType) (*entities.Token, error)
	SelectByAccessToken(accessToken string) (*entities.Token, error)
}
