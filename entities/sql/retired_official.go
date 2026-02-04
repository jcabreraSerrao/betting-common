package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type RetiredOfficial struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	RaceID             uint64           `gorm:"column:id_race;not null" json:"raceID"`
	Race               Race             `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID            uint64           `gorm:"column:group_id;not null" json:"groupID"`
	Group              Group            `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	ParticipantsRaceID uint             `gorm:"column:id_participants;not null" json:"participantsRaceID"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participantsRace"`
	Cancel             bool             `gorm:"column:cancel;default:false;not null" json:"cancel"`
	Reason             string           `gorm:"column:reason" json:"reason"`
	CreatedAt          time.Time        `json:"createdAt"`
	UpdatedAt          time.Time        `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (r *RetiredOfficial) TableName() string {
	return "gaming.retired_official"
}

func (r *RetiredOfficial) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RetiredOfficial
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

func (r *RetiredOfficial) UnmarshalJSON(data []byte) error {
	type Alias RetiredOfficial
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(data, aux)
}
