package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type ConfigRemate struct {
	RaceID        uint64             `gorm:"column:id_race" json:"idRace"`
	Race          Race               `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID       uint               `gorm:"column:id_group" json:"idGroup"`
	Group         Group              `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	Status        bool               `gorm:"column:status;default:false" json:"status"`
	CloseBoard    bool               `gorm:"column:close_board;default:false" json:"closeBoard"`
	TerciosRemate []TerciosRemate    `gorm:"foreignKey:IdConfigRemate;references:ID" json:"terciosRemate"`
	Ejemplares    []RemateEjemplares `gorm:"foreignKey:RemateID;references:ID" json:"ejemplares"`
	Transaction   []Transactions     `gorm:"foreignKey:RemateId;references:ID" json:"transactions"`
	ComboRemate   []ComboRemate      `gorm:"foreignKey:IDRemate;references:ID" json:"comboRemate"`
	gorm.Model
}

func (c *ConfigRemate) TableName() string {
	return "gaming.config_remate"
}

func (c *ConfigRemate) MarshalJSON() ([]byte, error) {
	if c.ID == 0 {
		return []byte("null"), nil
	}

	type Alias ConfigRemate
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(c),
		},
	)
}

func (c *ConfigRemate) UnmarshalJSON(data []byte) error {
	type Alias ConfigRemate
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
