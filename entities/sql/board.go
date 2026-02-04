package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Board struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	Position           int              `gorm:"column:position" json:"position"`
	Reverse            bool             `gorm:"column:reverse;default:false;not null" json:"reverse"`
	RaceID             uint64           `gorm:"column:id_race" json:"raceID"`
	Race               Race             `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	ParticipantsRaceID uint             `gorm:"column:id_participants" json:"idParticipants"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participants"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Board) TableName() string {
	return "gaming.board"
}

func (s *Board) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Board
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
