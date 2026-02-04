package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type RolesPermissions struct {
	IDRoles      uint64 `gorm:"column:id_roles"`
	IDPermission uint64 `gorm:"column:id_permission"`
	Status       bool   `gorm:"column:status"`

	// Relationships (using pointers for flexibility)
	Role       *Roles       `gorm:"foreignKey:IDRoles;references:ID"`
	Permission *Permissions `gorm:"foreignKey:IDPermission;references:ID"`
	gorm.Model
}

func (r *RolesPermissions) TableName() string {
	return "security.roles_permissions"
}

func (r *RolesPermissions) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RolesPermissions
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}
