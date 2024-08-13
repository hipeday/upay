package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryImpl struct {
	db *sqlx.DB
}

func (a *AccountRepositoryImpl) Setup(db *sqlx.DB) {
	a.db = db
}

func (a *AccountRepositoryImpl) TableName() string {
	return "account"
}

func (a *AccountRepositoryImpl) Columns() []string {
	return []string{"id", "username", "password", "email", "status", "secret", "create_at"}
}

func (a *AccountRepositoryImpl) GetDB() *sqlx.DB {
	return a.db
}

func (a *AccountRepositoryImpl) Columns2Query() string {
	columns := a.Columns()
	var columns2Query string
	for i, column := range columns {
		if i == 0 {
			columns2Query = column
		} else {
			columns2Query = fmt.Sprintf("%s, %s", columns2Query, column)
		}
	}
	return columns2Query
}

func (a *AccountRepositoryImpl) SelectAccountByUsername(username string) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (a *AccountRepositoryImpl) SelectSignIn(username string, password string) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ? AND password = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}
