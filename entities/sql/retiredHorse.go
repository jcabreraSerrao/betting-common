package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type RetiredHorse struct {
	RaceId             uint64           `gorm:"column:id_race" json:"idRace"`
	Race               Race             `gorm:"foreignKey:RaceId;references:ID"`
	Cancel             bool             `gorm:"column:cancel;default:false;not null" json:"cancel"`
	ParticipantsRaceID uint             `gorm:"column:id_participants" json:"idParticipants"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participants"`
	gorm.Model
}

func (s *RetiredHorse) TableName() string {
	return "gaming.retired_horse"
}

func (s *RetiredHorse) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RetiredHorse

	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
