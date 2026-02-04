package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name"`
	Level   int    `gorm:"column:level" json:"level"`
	GroupID int    `gorm:"column:id_group" json:"idGroup"`

	Group            Group              `gorm:"foreignKey:GroupID;references:ID;" json:"group"`
	RolesPermissions []RolesPermissions `gorm:"foreignKey:IDRoles;" json:"rolesPermissions"`
}

func (r *Roles) TableName() string {
	return "security.roles"
}

func (r *Roles) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Roles
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}
