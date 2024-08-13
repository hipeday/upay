package repository

import (
	"fmt"
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
	"strings"
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
}

const (
	insertSqlTemplate = "INSERT INTO %s (%s) VALUES (%s)"
)

func getInsertSql(tableName, columns string, columnCount int) string {
	var placeholder []string
	for i := 0; i < columnCount; i++ {
		placeholder = append(placeholder, "?")
	}
	return fmt.Sprintf(insertSqlTemplate, tableName, columns, strings.Join(placeholder, ","))
}
