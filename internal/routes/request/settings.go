package request

import "github.com/hipeday/upay/internal/entities"

type SaveSettingsPayload struct {
	Config      string                `json:"config" validate:"required" message:"Config is illegal"`
	Name        *string               `json:"name"`
	Value       *string               `json:"value"`
	Type        entities.SettingsType `json:"type"`
	Required    bool                  `json:"required"`
	Description *string               `json:"description"`
	OperatorId  int64                 `json:"operator_id"`
}

func (s SaveSettingsPayload) Validate() error {
	return Validate(s)
}
