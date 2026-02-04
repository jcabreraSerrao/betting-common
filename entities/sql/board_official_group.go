package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BoardOfficialGroup struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	Position           int              `gorm:"column:position;not null" json:"position"`
	Reverse            bool             `gorm:"column:reverse;default:false;not null" json:"reverse"`
	RaceID             uint64           `gorm:"column:id_race;not null" json:"raceID"`
	Race               Race             `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID            uint64           `gorm:"column:group_id;not null" json:"groupID"`
	Group              Group            `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	ParticipantsRaceID uint             `gorm:"column:id_participants;not null" json:"participantsRaceID"`
	ParticipantsRace   ParticipantsRace `gorm:"foreignKey:ParticipantsRaceID;references:ID" json:"participantsRace"`
	CreatedAt          time.Time        `json:"createdAt"`
	UpdatedAt          time.Time        `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (b *BoardOfficialGroup) TableName() string {
	return "gaming.board_official_group"
}

func (b *BoardOfficialGroup) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias BoardOfficialGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
		},
	)
}

func (b *BoardOfficialGroup) UnmarshalJSON(data []byte) error {
	type Alias BoardOfficialGroup
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}
