package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type SubGroup struct {
	Name    string `gorm:"column:name"`
	GroupID uint   `gorm:"column:id_group"`

	Group *Group `gorm:"foreignKey:GroupID;references:ID"`
	gorm.Model
}

func (s *SubGroup) TableName() string {
	return "security.sub_group"
}

func (s *SubGroup) MarshalJSON() ([]byte, error) {
	type Alias SubGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
