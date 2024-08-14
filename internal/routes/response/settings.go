package response

import "github.com/hipeday/upay/internal/entities"

type SaveSettings struct {
	Config      string                `json:"config"`
	Name        *string               `json:"name"`
	Value       *string               `json:"value"`
	Type        entities.SettingsType `json:"type"`
	Required    bool                  `json:"required"`
	Description *string               `json:"description"`
}
