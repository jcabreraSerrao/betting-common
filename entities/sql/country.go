package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Country represents a country in the system.
// Contains information about the country name, code, and associated currency.
type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`      // Country name (Venezuela, United States)
	Code string `gorm:"column:code;type:varchar(3);unique;not null" json:"code"` // Country code (VE, US, CO)
	// CurrencyID is nullable to allow ON DELETE SET NULL behavior when a currency is removed.
	CurrencyID *uint          `gorm:"column:currency_id" json:"currency_id"`
	Currency   *Currency      `gorm:"foreignKey:CurrencyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"currency"` // Associated currency
	Status     bool           `gorm:"column:status;type:boolean;default:true" json:"status"`                                              // Active/inactive status
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
	// Hippodromes is the reverse relation for hipodromos in this country.
	Hippodromes               []Hippodromes              `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"hipodromos,omitempty"`
	GroupCountryMinBetConfigs []GroupCountryMinBetConfig `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"groupCountryMinBetConfigs,omitempty"`
}

// TableName returns the table name for the Country entity.
// This method is used by GORM to map the structure to the correct table.
func (c *Country) TableName() string {
	return "config.country"
}

// MarshalJSON customizes the JSON serialization of the Country entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (c *Country) MarshalJSON() ([]byte, error) {
	if c.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Country
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(c),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the Country entity.
func (c *Country) UnmarshalJSON(data []byte) error {
	type Alias Country
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
