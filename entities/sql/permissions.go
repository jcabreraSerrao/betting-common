package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Permissions struct {
	Key   string `gorm:"column:key" json:"key"`
	Alias string `gorm:"column:alias" json:"alias"`
	Group string `gorm:"column:group" json:"group"`
	gorm.Model
}

func (p *Permissions) TableName() string {
	return "security.permissions"
}

func (p *Permissions) MarshalJSON() ([]byte, error) {
	if p.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Permissions
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(p),
		},
	)
}
