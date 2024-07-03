package repository

import "github.com/jmoiron/sqlx"

type AccountRepository struct {
	db *sqlx.DB
}

func (a AccountRepository) Setup(db *sqlx.DB) {
	a.db = db
}
