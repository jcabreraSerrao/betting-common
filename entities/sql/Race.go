package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Race struct {
	ID                   uint64                `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                 string                `gorm:"column:name" json:"name"`
	Open                 bool                  `gorm:"column:open;default:true;not null" json:"open"`
	OpenDate             time.Time             `gorm:"column:open_date" json:"openDate"`
	Distance             string                `gorm:"column:distance" json:"distance"`
	Bet                  []Bet                 `gorm:"foreignKey:RaceID;references:ID" json:"bet"`
	Board                []*Board              `gorm:"foreignKey:RaceID;references:ID" json:"board"`
	RetiredHorse         []*RetiredHorse       `gorm:"foreignKey:RaceId;references:ID" json:"retiredHorse"`
	BoardRaceParada      []*BoardRaceParada    `gorm:"foreignKey:RaceID;references:ID" json:"boardRaceParada"`
	RetiredHorseGroup    []*RetiredHorseGroup  `gorm:"foreignKey:RaceId;references:ID" json:"retiredHorseGroup"`
	ConfigRemate         []*ConfigRemate       `gorm:"foreignKey:RaceID;references:ID" json:"configRemate"`
	ParticipantsRace     []ParticipantsRace    `gorm:"foreignKey:RaceID;references:ID" json:"participants"`
	RacesProcessGroups   []RacesProcessGroup   `gorm:"foreignKey:RaceID;references:ID" json:"racesProcessGroups"`
	RaceGroupCommissions []RaceGroupCommission `gorm:"foreignKey:RaceID;references:ID" json:"raceGroupCommissions"`
	BoardOfficialGroups  []BoardOfficialGroup  `gorm:"foreignKey:RaceID;references:ID" json:"boardOfficialGroups"`
	RaceDividendGroups   []RaceDividendGroup   `gorm:"foreignKey:RaceID;references:ID" json:"raceDividendGroups"`
	RaceDividends        []RaceDividend        `gorm:"foreignKey:RaceID;references:ID" json:"raceDividends"`
	RetiredOfficial      []RetiredOfficial     `gorm:"foreignKey:RaceID;references:ID" json:"retiredOfficial"`
	HippodromeId         uint                  `gorm:"column:hippodrome_group" json:"hippodromeId"`
	Hippodrome           Hippodromes           `gorm:"foreignKey:HippodromeId;references:ID" json:"hippodrome"`
	IsCancelled          bool                  `gorm:"column:is_cancelled;default:false" json:"isCancelled"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Race) TableName() string {
	return "gaming.race"
}

func (s *Race) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Race
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}

func (s *Race) UnmarshalJSON(data []byte) error {
	type Alias Race
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
