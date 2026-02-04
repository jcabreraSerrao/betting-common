package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BoardRaceParada struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	Position           int              `gorm:"column:position" json:"position"`
	Reverse            bool             `gorm:"column:reverse;default:false;not null" json:"reverse"`
	RaceID             uint64           `gorm:"column:id_race" json:"raceID"`
	Race               Race             `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID            uint64           `gorm:"column:id_group" json:"groupID"`
	Group              Group            `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	ParticipantsRaceID uint             `gorm:"column:id_participants" json:"idParticipants"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participants"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *BoardRaceParada) TableName() string {
	return "gaming.board_race_group"
}

func (s *BoardRaceParada) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias BoardRaceParada
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
