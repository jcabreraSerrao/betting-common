package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type RetiredHorseGroup struct {
	RaceId             uint64           `gorm:"column:id_race" json:"idRace"`
	Race               Race             `gorm:"foreignKey:RaceId;references:ID"`
	GroupID            uint64           `gorm:"column:id_group;not null" json:"idGroup"`
	Cancel             bool             `gorm:"column:cancel;default:false;not null" json:"cancel"`
	ParticipantsRaceID uint             `gorm:"column:id_participants" json:"idParticipants"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participants"`
	gorm.Model
}

func (s *RetiredHorseGroup) TableName() string {
	return "gaming.retired_horse_group"
}

func (s *RetiredHorseGroup) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RetiredHorseGroup

	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
