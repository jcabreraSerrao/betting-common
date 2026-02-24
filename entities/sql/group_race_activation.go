package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// GroupRaceActivation representa la activación de una carrera para un grupo específico.
type GroupRaceActivation struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID        uint64         `gorm:"column:group_id;not null;index:idx_group_race" json:"group_id"`
	Group          Group          `gorm:"foreignKey:GroupID" json:"-"`
	RaceID         uint64         `gorm:"column:race_id;not null;index:idx_group_race" json:"race_id"`
	Race           Race           `gorm:"foreignKey:RaceID" json:"-"`
	IsActive       bool           `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	LastNotifiedAt *time.Time     `gorm:"column:last_notified_at" json:"last_notified_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (GroupRaceActivation) TableName() string {
	return "gaming.group_race_activations"
}

func (g *GroupRaceActivation) MarshalJSON() ([]byte, error) {
	type Alias GroupRaceActivation
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(g),
	})
}
