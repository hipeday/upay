package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
	"strings"
)

type AccountRepositoryImpl struct {
	AccountRepository
	db *sqlx.DB
}

func (a *AccountRepositoryImpl) Setup(db *sqlx.DB) {
	a.db = db
}

func (a *AccountRepositoryImpl) TableName() string {
	return "account"
}

func (a *AccountRepositoryImpl) Columns() []string {
	return getColumns(entities.Account{})
}

func (a *AccountRepositoryImpl) GetDB() *sqlx.DB {
	return a.db
}

func (a *AccountRepositoryImpl) Columns2Query() string {
	return strings.Join(a.Columns(), ", ")
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
