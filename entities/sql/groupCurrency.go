package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// GroupCurrency represents the currency configuration for a specific group.
// Contains information about which currencies a group can use and their settings.
type GroupCurrency struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID    uint           `gorm:"column:group_id;not null" json:"group_id"`                 // Foreign key to Group
	Group      Group          `gorm:"foreignKey:GroupID;references:ID" json:"group"`            // Associated group
	CurrencyID uint           `gorm:"column:currency_id;not null" json:"currency_id"`           // Foreign key to Currency
	Currency   Currency       `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`      // Associated currency
	IsMain     bool           `gorm:"column:is_main;type:boolean;default:false" json:"is_main"` // Indicates if this is the main display currency for the group
	CustomRate *float64       `gorm:"column:custom_rate;type:decimal(15,6)" json:"custom_rate"` // Optional custom exchange rate for this group
	Status     bool           `gorm:"column:status;type:boolean;default:true" json:"status"`    // Active/inactive status
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the GroupCurrency entity.
// This method is used by GORM to map the structure to the correct table.
func (gc *GroupCurrency) TableName() string {
	return "config.group_currency"
}

// MarshalJSON customizes the JSON serialization of the GroupCurrency entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (gc *GroupCurrency) MarshalJSON() ([]byte, error) {
	if gc.ID == 0 {
		return []byte("null"), nil
	}

	type Alias GroupCurrency
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(gc),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the GroupCurrency entity.
func (gc *GroupCurrency) UnmarshalJSON(data []byte) error {
	type Alias GroupCurrency
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(gc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
