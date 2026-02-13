package sql

import (
	"time"

	"gorm.io/gorm"
)

type PollaRace struct {
	ID            uint64              `gorm:"primaryKey;autoIncrement" json:"id"`
	PollaID       uint64              `gorm:"column:polla_id;index;not null" json:"pollaId"`
	Polla         Polla               `gorm:"foreignKey:PollaID;references:ID" json:"polla"`
	RaceID        uint64              `gorm:"column:id_race;index;not null" json:"raceID"`
	Race          Race                `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	RaceOrder     int                 `gorm:"column:race_order;not null" json:"raceOrder"` // Orden de la carrera en la polla (1-6)
	InvalidHorses []PollaInvalidHorse `gorm:"foreignKey:PollaRaceID;references:ID" json:"invalidHorses"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PollaRace) TableName() string {
	return "gaming.polla_race"
}
