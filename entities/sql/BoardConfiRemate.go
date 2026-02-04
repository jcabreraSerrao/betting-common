package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BoardConfiRemate struct {
	ID             uint64         `gorm:"primarykey;autoIncrement" json:"id"`
	Reverse        bool           `gorm:"column:reverse;default:false;not null" json:"reverse"`
	BoardID        uint           `gorm:"column:id_board" json:"idBoard"`
	Board          Board          `gorm:"foreignKey:BoardID" json:"board"`
	RaceID         uint64         `gorm:"column:id_race" json:"idRace"`
	Race           Race           `gorm:"foreignKey:RaceID" json:"race"`
	ConfigRemateID uint           `gorm:"column:id_config_remate" json:"idConfigRemate"`
	ConfigRemate   ConfigRemate   `gorm:"foreignKey:ConfigRemateID" json:"configRemate"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

func (b *BoardConfiRemate) TableName() string {
	return "gaming.board_config_remate"
}

func (b *BoardConfiRemate) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias BoardConfiRemate
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
		},
	)
}

func (b *BoardConfiRemate) UnmarshalJSON(data []byte) error {
	type Alias BoardConfiRemate
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
