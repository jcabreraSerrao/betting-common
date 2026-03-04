package sql

import (
	"time"

	"gorm.io/gorm"
)

type PollaInvalidHorse struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	PollaRaceID uint64 `gorm:"column:polla_race_id;index;not null" json:"pollaRaceId"`
	Position    string `gorm:"column:position;not null" json:"position"` // Posición del caballo invalidado (ej: "1", "1a")
	Reason      string `gorm:"column:reason;type:varchar(255)" json:"reason"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PollaInvalidHorse) TableName() string {
	return "gaming.polla_invalid_horse"
}
