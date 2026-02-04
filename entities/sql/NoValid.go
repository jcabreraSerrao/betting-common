package sql

import (
	"time"

	"gorm.io/gorm"
)

type NoValid struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	BetID              uint64           `gorm:"column:bet_id;not null" json:"betId"`
	Bet                Bet              `gorm:"foreignKey:BetID;references:ID" json:"bet"`
	ParticipantsRaceID uint             `gorm:"column:participants_race_id;not null" json:"participantsRaceId"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participantsRace"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (n *NoValid) TableName() string {
	return "gaming.no_valid"
}
