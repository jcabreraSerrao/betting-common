package sql

import (
	"time"

	"gorm.io/gorm"
)

type PollaSelection struct {
	ID                 uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	PollaParticipantID uint64 `gorm:"column:polla_participant_id;index;not null" json:"pollaParticipantId"`
	PollaRaceID        uint64 `gorm:"column:polla_race_id;index;not null" json:"pollaRaceId"`
	ParticipantRaceID  uint64 `gorm:"column:participant_race_id;index;not null" json:"participantRaceId"` // FK a ParticipantsRace (Gaming)
	Points             int    `gorm:"column:points;default:0" json:"points"`                              // Calculado al cerrar (0, 1, 3, 5)
	ActualPosition     *int   `gorm:"column:actual_position" json:"actualPosition"`                       // Posición final ajustada tras descartar invalidados
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PollaSelection) TableName() string {
	return "gaming.polla_selection"
}
