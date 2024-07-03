package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func (a AccountRepository) Setup(db *sqlx.DB) {
	a.db = db
}

func (a AccountRepository) TableName() string {
	return "account"
}

func (a AccountRepository) Columns() []string {
	return []string{"id", "username", "password", "status", "token", "refresh_token", "create_at"}
}

func (a AccountRepository) Columns2Query() string {
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
