package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// PaymentPlatformCurrency represents the association between payment platforms and currencies.
// Allows tracking which currencies each payment platform supports (fiat and crypto).
type PaymentPlatformCurrency struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	PaymentPlatformID uint            `gorm:"column:payment_platform_id;not null" json:"payment_platform_id"`     // Foreign key to PaymentPlatform
	PaymentPlatform   PaymentPlatform `gorm:"foreignKey:PaymentPlatformID;references:ID" json:"payment_platform"` // Associated payment platform
	CurrencyID        uint            `gorm:"column:currency_id;not null" json:"currency_id"`                     // Foreign key to Currency
	Currency          Currency        `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`                // Associated currency
	Address           string          `gorm:"column:address;type:varchar(255)" json:"address"`                    // Address for wallets/exchanges
	Status            bool            `gorm:"column:status;type:boolean;default:true" json:"status"`              // Active/inactive status
	CreatedAt         time.Time       `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt         gorm.DeletedAt  `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the PaymentPlatformCurrency entity.
// This method is used by GORM to map the structure to the correct table.
func (ppc *PaymentPlatformCurrency) TableName() string {
	return "config.payment_platform_currency"
}

// MarshalJSON customizes the JSON serialization of the PaymentPlatformCurrency entity.
// If the ID is 0, returns "null". Otherwise, serializes the complete structure.
func (ppc *PaymentPlatformCurrency) MarshalJSON() ([]byte, error) {
	if ppc.ID == 0 {
		return []byte("null"), nil
	}

	type Alias PaymentPlatformCurrency
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(ppc),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the PaymentPlatformCurrency entity.
func (ppc *PaymentPlatformCurrency) UnmarshalJSON(data []byte) error {
	type Alias PaymentPlatformCurrency
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(ppc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
