package sql

import (
	"time"

	"gorm.io/gorm"
)

type ParticipantsRace struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	RaceID        uint           `gorm:"column:race_id" json:"idRace"`
	Race          Race           `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	Position      int            `gorm:"column:position" json:"position"`
	Horse         string         `gorm:"column:horse" json:"horse"`
	Jockey        string         `gorm:"column:jockey" json:"jockey"`
	Trainer       string         `gorm:"column:trainer" json:"trainer"`
	Retired       bool           `gorm:"column:retired;default:false" json:"retired"`
	TerciosRemate *TerciosRemate `gorm:"foreignKey:ParticipantsRaceId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_gaming_participants" json:"terciosRemate"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *ParticipantsRace) TableName() string {
	return "gaming.participants_race"
}
