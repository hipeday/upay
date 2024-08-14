package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
)

type TokenRepositoryImpl struct {
	db *sqlx.DB
}

func (t *TokenRepositoryImpl) Setup(db *sqlx.DB) {
	t.db = db
}

func (t *TokenRepositoryImpl) TableName() string {
	return "token"
}

func (t *TokenRepositoryImpl) Columns() []string {
	return []string{"id", "target_id", "type", "access_token", "refresh_token", "expires_at", "create_at"}
}

func (t *TokenRepositoryImpl) GetDB() *sqlx.DB {
	return t.db
}

func (t *TokenRepositoryImpl) Columns2Query() string {
	columns := t.Columns()
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

func (t *TokenRepositoryImpl) Insert(token *entities.Token) error {
	db := t.db
	tx := db.MustBegin()
	tx.MustExec(getInsertSql(t.TableName(), t.Columns2Query(), len(t.Columns())), nil, token.TargetId, token.Type, token.AccessToken, token.RefreshToken, token.ExpiresAt, token.CreateAt)
	return tx.Commit()
}

func (t *TokenRepositoryImpl) SelectByTargetId(targetId int64, tokenType entities.TokenType) (*entities.Token, error) {
	db := t.db
	var token entities.Token
	err := db.Get(&token, getQuerySql(t.TableName(), t.Columns2Query(), []string{"target_id", "type"}), targetId, tokenType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}

func (t *TokenRepositoryImpl) UpdateById(token *entities.Token) error {
	db := t.db
	tx := db.MustBegin()
	tx.MustExec(getUpdateSql(t.TableName(), []string{"target_id", "type", "access_token", "refresh_token", "expires_at", "create_at"}, []string{"id"}), token.TargetId, token.Type, token.AccessToken, token.RefreshToken, token.ExpiresAt, token.CreateAt, token.ID)
	return tx.Commit()
}

func (t *TokenRepositoryImpl) SelectByAccessToken(accessToken string) (*entities.Token, error) {
	db := t.db
	var token entities.Token
	err := db.Get(&token, getQuerySql(t.TableName(), t.Columns2Query(), []string{"access_token"}), accessToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}
