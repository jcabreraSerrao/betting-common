package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// GroupExchangeRate represents exchange rates specific to a group.
// Allows groups to have their own exchange rates with audit trail and banca propagation.
type GroupExchangeRate struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID        uint           `gorm:"column:group_id;not null" json:"group_id"`                                      // Foreign key to Group
	Group          Group          `gorm:"foreignKey:GroupID;references:ID" json:"group"`                                 // Associated group
	FromCurrencyID uint           `gorm:"column:from_currency_id;not null" json:"from_currency_id"`                      // Source currency ID
	FromCurrency   Currency       `gorm:"foreignKey:FromCurrencyID;references:ID" json:"from_currency"`                  // Source currency relation
	ToCurrencyID   uint           `gorm:"column:to_currency_id;not null" json:"to_currency_id"`                          // Target currency ID (USD or group's base)
	ToCurrency     Currency       `gorm:"foreignKey:ToCurrencyID;references:ID" json:"to_currency"`                      // Target currency relation
	Rate           float64        `gorm:"column:rate;type:decimal(15,6);not null" json:"rate"`                           // Exchange rate (e.g., 150.0 for VES→USD)
	Source         string         `gorm:"column:source;type:varchar(255)" json:"source"`                                 // Rate source (BCV, Market, Manual)
	LoadedBy       *uint          `gorm:"column:loaded_by" json:"loaded_by"`                                             // User who loaded the rate
	LoadedAt       time.Time      `gorm:"column:loaded_at;type:timestamp with time zone;default:now()" json:"loaded_at"` // When the rate was loaded
	Status         bool           `gorm:"column:status;type:boolean;default:true" json:"status"`                         // Active/inactive status
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt      gorm.DeletedAt `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the GroupExchangeRate entity.
// This method is used by GORM to map the structure to the correct table.
func (ger *GroupExchangeRate) TableName() string {
	return "config.group_exchange_rate"
}

// MarshalJSON customizes the JSON serialization of the GroupExchangeRate entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (ger *GroupExchangeRate) MarshalJSON() ([]byte, error) {
	if ger.ID == 0 {
		return []byte("null"), nil
	}

	type Alias GroupExchangeRate
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(ger),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the GroupExchangeRate entity.
func (ger *GroupExchangeRate) UnmarshalJSON(data []byte) error {
	type Alias GroupExchangeRate
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(ger),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
