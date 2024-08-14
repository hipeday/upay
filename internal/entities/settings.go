package entities

type SettingsType string

const (
	StringSettingsType = "string"
)

type Settings struct {
	Entity
	Config      string       `db:"config"`
	Name        *string      `db:"name"`
	Value       *string      `db:"value"`
	Required    bool         `db:"required"`
	Type        SettingsType `db:"type"`
	Description *string      `db:"description"`
	ModifiedBy  int64        `db:"modified_by"`
}
