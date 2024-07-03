package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Setup(db *sqlx.DB)
}
