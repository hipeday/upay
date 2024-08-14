package repository

import (
	"database/sql"
	"errors"
	"github.com/hipeday/upay/internal/entities"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SettingsRepositoryImpl struct {
	SettingsRepository
	db *sqlx.DB
}

func (s *SettingsRepositoryImpl) Setup(db *sqlx.DB) {
	s.db = db
}

func (s *SettingsRepositoryImpl) TableName() string {
	return "settings"
}

func (s *SettingsRepositoryImpl) Columns() []string {
	return getColumns(entities.Settings{})
}

func (s *SettingsRepositoryImpl) Columns2Query() string {
	return strings.Join(s.Columns(), ", ")
}

func (s *SettingsRepositoryImpl) GetDB() *sqlx.DB {
	return s.db
}

func (s *SettingsRepositoryImpl) Insert(settings entities.Settings) error {
	db := s.db
	tx := db.MustBegin()
	tx.MustExec(getInsertSql(s.TableName(), s.Columns2Query(), len(s.Columns())), nil, settings.CreateAt, settings.Config, settings.Name, settings.Value, settings.Required, settings.Type, settings.Description, settings.ModifiedBy)
	return tx.Commit()
}

func (s *SettingsRepositoryImpl) UpdateById(settings *entities.Settings) error {
	db := s.db
	tx := db.MustBegin()
	tx.MustExec(getUpdateSql(s.TableName(), []string{"config", "name", "value", "required", "type", "description", "modified_by", "create_at"}, []string{"id"}), settings.Config, settings.Name, settings.Value, settings.Required, settings.Type, settings.Description, settings.ModifiedBy, settings.CreateAt, settings.ID)
	return tx.Commit()
}

func (s *SettingsRepositoryImpl) SelectByConfig(configKey string) (*entities.Settings, error) {
	db := s.db
	var settings entities.Settings
	err := db.Get(&settings, getQuerySql(s.TableName(), s.Columns2Query(), []string{"config"}), configKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &settings, nil
}
