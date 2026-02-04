package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Hippodromes struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Slug string `gorm:"column:slug" json:"slug"`
	// CountryID is optional; when a country is deleted, set to NULL.
	CountryID *uint    `gorm:"column:country_id" json:"country_id"`
	Country   *Country `gorm:"foreignKey:CountryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country"` // Associated country
	Status    bool     `gorm:"column:status" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (h *Hippodromes) TableName() string {
	return "config.hipodromos"
}

func (h *Hippodromes) MarshalJSON() ([]byte, error) {
	if h.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Hippodromes
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(h),
		},
	)
}
