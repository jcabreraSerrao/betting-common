package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Currency represents a currency in the system (fiat or crypto).
// Contains information about the currency code, name, symbol, and type.
type Currency struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string         `gorm:"column:code;type:varchar(10);unique;not null" json:"code"`     // Currency code (USD, VES, BTC, ETH)
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`           // Currency name (US Dollar, Venezuelan Bolívar)
	Symbol    string         `gorm:"column:symbol;type:varchar(10)" json:"symbol"`                 // Currency symbol ($, Bs., ₿)
	IsCrypto  bool           `gorm:"column:is_crypto;type:boolean;default:false" json:"is_crypto"` // Indicates if it's a cryptocurrency
	Status    bool           `gorm:"column:status;type:boolean;default:true" json:"status"`        // Active/inactive status
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
	// Countries is the reverse relation for currencies used by countries.
	Countries []Country `gorm:"foreignKey:CurrencyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"countries,omitempty"`
}

// TableName returns the table name for the Currency entity.
// This method is used by GORM to map the structure to the correct table.
func (c *Currency) TableName() string {
	return "config.currency"
}

// MarshalJSON customizes the JSON serialization of the Currency entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (c *Currency) MarshalJSON() ([]byte, error) {
	if c.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Currency
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(c),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the Currency entity.
func (c *Currency) UnmarshalJSON(data []byte) error {
	type Alias Currency
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
