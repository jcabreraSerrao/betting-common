package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type TypeBet struct {
	Name           string         `gorm:"column:name" json:"name"`
	TypeCalc       string         `gorm:"column:type_calc" json:"type_calc"`
	Classification string         `gorm:"column:classification" json:"classification"`
	TypeBetGroup   []TypeBetGroup `gorm:"foreignKey:TypeBetId;references:ID" json:"typeBetGroup"`

	gorm.Model
}

func (s *TypeBet) TableName() string {
	return "gaming.type_bet"
}

func (s *TypeBet) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TypeBet
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
