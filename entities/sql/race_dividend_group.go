package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type RaceDividendGroup struct {
	ID            uint             `gorm:"primarykey" json:"id"`
	RaceID        uint64           `gorm:"column:id_race;not null" json:"raceID"`
	Race          Race             `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID       uint64           `gorm:"column:group_id;not null" json:"groupID"`
	Group         Group            `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	ParticipantID uint             `gorm:"column:participant_id;not null" json:"participantID"`
	Participant   ParticipantsRace `gorm:"foreignKey:ParticipantID;references:ID" json:"participant"`
	TypeBetID     uint             `gorm:"column:type_bet_id;not null" json:"typeBetID"`
	TypeBet       TypeBet          `gorm:"foreignKey:TypeBetID;references:ID" json:"typeBet"`
	Dividend      float64          `gorm:"column:dividend;not null" json:"dividend"`
	LoadedBy      *uint            `gorm:"column:loaded_by" json:"loadedBy"`
	LoadedAt      *time.Time       `gorm:"column:loaded_at" json:"loadedAt"`
	Source        string           `gorm:"column:source" json:"source"`
	Status        bool             `gorm:"column:status;default:true;not null" json:"status"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (r *RaceDividendGroup) TableName() string {
	return "gaming.race_dividend_group"
}

func (r *RaceDividendGroup) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceDividendGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

func (r *RaceDividendGroup) UnmarshalJSON(data []byte) error {
	type Alias RaceDividendGroup
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(data, aux)
}
