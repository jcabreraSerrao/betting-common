package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type TypeBetGroup struct {
	Alias     string  `gorm:"column:alias" json:"alias"`
	TypeBetId uint    `gorm:"column:id_type_bet" json:"id_type_bet"`
	GroupId   uint    `gorm:"column:id_group" json:"id_group"`
	Status    bool    `gorm:"column:status" json:"status"`
	Group     Group   `gorm:"foreignKey:GroupId;references:ID" json:"group"`
	TypeBet   TypeBet `gorm:"foreignKey:TypeBetId;references:ID" json:"typeBet"`

	gorm.Model
}

func (s *TypeBetGroup) TableName() string {
	return "gaming.type_bet_group"
}

func (s *TypeBetGroup) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TypeBetGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
