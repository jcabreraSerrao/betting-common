package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// GroupCountryMinBetConfig represents minimum bet configuration per country/group.
// This entity manages minimum bet amounts by country, and optionally by group,
// allowing banks to create/edit configurations for their associated groups.
// The currency is obtained from the country.
type GroupCountryMinBetConfig struct {
	ID                 uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CountryID          uint            `gorm:"column:country_id;not null" json:"country_id"`
	Country            Country         `gorm:"foreignKey:CountryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"country"`
	GroupID            *uint           `gorm:"column:group_id" json:"group_id"`                                                                       // Optional group override
	Group              *Group          `gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group,omitempty"` // Associated group
	MinBetAmount       decimal.Decimal `gorm:"column:min_bet_amount;type:decimal(18,8);not null" json:"min_bet_amount"`                               // Minimum bet amount
	Percentage         int             `gorm:"column:percentage;default:100" json:"percentage"`                                                       // Percentage (1-100)
	AllowGroupOverride bool            `gorm:"column:allow_group_override;default:false" json:"allow_group_override"`                                 // Whether group can create its own override
	IsStrictOverride   bool            `gorm:"column:is_strict_override;default:false" json:"is_strict_override"`                                     // Forces priority if applicable
	Active             bool            `gorm:"column:active;default:true" json:"active"`                                                              // Active/inactive status
	CreatedBy          *uint           `gorm:"column:created_by" json:"created_by"`                                                                   // Who created this configuration
	CreatedAt          time.Time       `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt          time.Time       `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt          gorm.DeletedAt  `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the GroupCountryMinBetConfig entity.
// This method is used by GORM to map the structure to the correct table.
func (g *GroupCountryMinBetConfig) TableName() string {
	return "config.group_country_min_bet_config"
}

// MarshalJSON customizes the JSON serialization of the GroupCountryMinBetConfig entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (g *GroupCountryMinBetConfig) MarshalJSON() ([]byte, error) {
	if g.ID == 0 {
		return []byte("null"), nil
	}

	type Alias GroupCountryMinBetConfig
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(g),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the GroupCountryMinBetConfig entity.
func (g *GroupCountryMinBetConfig) UnmarshalJSON(data []byte) error {
	type Alias GroupCountryMinBetConfig
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(g),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
