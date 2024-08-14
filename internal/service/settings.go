package service

import (
	"fmt"
	"github.com/hipeday/upay/internal/entities"
	errors2 "github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/routes/request"
	"github.com/hipeday/upay/internal/routes/response"
	"time"
)

type SettingsServiceImpl struct {
	SettingsService
	repository repository.SettingsRepository
}

func (s *SettingsServiceImpl) Setup(repository repository.SettingsRepository) {
	s.repository = repository
}

func (s *SettingsServiceImpl) Save(payload request.SaveSettingsPayload) (*response.SaveSettings, error) {
	settingsRepository := s.repository
	settings, err := settingsRepository.SelectByConfig(payload.Config)
	if err != nil {
		return nil, err
	}
	if settings != nil {
		return nil, errors2.NewConflictErrorError(fmt.Sprintf("settings %s already exists", payload.Config))
	}
	now := time.Now()
	savedSettings := entities.Settings{
		Entity: entities.Entity{
			ID:       nil,
			CreateAt: &now,
		},
		Config:      payload.Config,
		Name:        payload.Name,
		Value:       payload.Value,
		Required:    payload.Required,
		Type:        payload.Type,
		Description: payload.Description,
		ModifiedBy:  payload.OperatorId,
	}

	err = settingsRepository.Insert(savedSettings)
	if err != nil {
		return nil, err
	}
	return &response.SaveSettings{
		Config:      savedSettings.Config,
		Name:        savedSettings.Name,
		Value:       savedSettings.Value,
		Type:        savedSettings.Type,
		Required:    savedSettings.Required,
		Description: savedSettings.Description,
	}, nil
}
