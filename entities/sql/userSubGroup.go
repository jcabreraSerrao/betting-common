package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type UserSubGroup struct {
	SubGroupID uint       `gorm:"column:id_sub_group"`
	UserID     uint       `gorm:"column:id_user_group"`
	Status     bool       `gorm:"column:status"`
	SubGroup   *SubGroup  `gorm:"foreignKey:SubGroupID;references:ID"`
	User       *UserGroup `gorm:"foreignKey:UserID;references:ID"`

	gorm.Model
}

func (s *UserSubGroup) TableName() string {
	return "security.user_sub_group"
}

func (s *UserSubGroup) MarshalJSON() ([]byte, error) {
	type Alias UserSubGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
