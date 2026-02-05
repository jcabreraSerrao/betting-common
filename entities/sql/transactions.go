package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Transactions representa una transacción en el sistema.
// Contiene información sobre el monto, comisión, roles, y relaciones con otras entidades.
type Transactions struct {
	ID                 uint64             `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	TerciosId          uint64             `gorm:"column:id_tercios" json:"terciosId"`
	Amount             decimal.Decimal    `gorm:"column:amount;type:decimal(18,8)" json:"amount"`         // Tipo ajustado para moneda
	Commission         decimal.Decimal    `gorm:"column:commission;type:decimal(18,8)" json:"commission"` // Tipo ajustado para moneda
	Rever              bool               `gorm:"column:rever;type:boolean;default:false" json:"rever"`   // Tipo booleano explícito
	Rol                string             `gorm:"column:rol;type:varchar(255)" json:"rol"`                // Tipo varchar explícito
	TransactionsId     *uint64            `gorm:"column:id_transaction" json:"transactionsId"`
	Transactions       *Transactions      `gorm:"foreignKey:TransactionsId;references:ID"` // Eliminada etiqueta de restricción explícita
	BetsId             *uint64            `gorm:"column:id_bet" json:"id_bet"`
	Bet                Bet                `gorm:"foreignKey:BetsId;references:ID"`
	TypeTransactionId  uint               `gorm:"column:id_type_transaction" json:"typeTransactionId"`
	TypeTransaction    TypeTransaction    `gorm:"foreignKey:TypeTransactionId;references:ID" json:"typeTransaction"`
	RaceId             *uint64            `gorm:"column:id_race" json:"id_race"`
	Race               Race               `gorm:"foreignKey:RaceId;references:ID"`
	DateTransaction    time.Time          `gorm:"column:date_transaction;type:timestamp with time zone;default:now();not null" json:"dateTransaction"` // Tipo timestamp explícito
	RemateId           *uint              `gorm:"column:id_remate" json:"id_remate"`
	Remate             ConfigRemate       `gorm:"foreignKey:RemateId;references:ID"`
	Tercios            Tercios            `gorm:"foreignKey:TerciosId;references:ID" json:"tercio"`
	Refills            *Refills           `gorm:"foreignKey:TransactionId;references:ID"` // TransactionId en Refills
	Withdrawals        *Withdrawal        `gorm:"foreignKey:IdTransaction;references:ID"` // Corregido a IdTransaction para coincidir con Withdrawal
	TercioRemate       *TerciosRemate     `gorm:"foreignKey:IdTransaction;references:ID"`
	CurrencyID         *uint              `gorm:"column:currency_id" json:"currency_id"`                                                                           // Foreign key to Currency
	Currency           *Currency          `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`                                                             // Associated currency
	ExchangeRate       *float64           `gorm:"column:exchange_rate;type:decimal(15,6)" json:"exchange_rate"`                                                    // Exchange rate applied at transaction time
	AmountUSD          *decimal.Decimal   `gorm:"column:amount_usd;type:decimal(15,6)" json:"amount_usd"`                                                          // Amount converted to USD for auditing
	AmountOriginal     decimal.Decimal    `gorm:"column:amount_original;type:decimal(18,8)" json:"amount_original"`                                                // Original amount before conversion
	CurrencyOriginalID *uint              `gorm:"column:currency_original_id" json:"currency_original_id"`                                                         // Original currency ID
	CurrencyOriginal   Currency           `gorm:"foreignKey:CurrencyOriginalID;references:ID" json:"currency_original"`                                            // Original currency relation
	ExchangeRateID     *uint              `gorm:"column:exchange_rate_id" json:"exchange_rate_id"`                                                                 // ID of the GroupExchangeRate used
	ExchangeRateRef    *GroupExchangeRate `gorm:"foreignKey:ExchangeRateID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"exchange_rate_ref"` // GroupExchangeRate relation
	CreatedAt          time.Time          `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt          time.Time          `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt          gorm.DeletedAt     `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName devuelve el nombre de la tabla en la base de datos para la entidad Transactions.
// Este método es utilizado por GORM para mapear la estructura a la tabla correcta.
func (s *Transactions) TableName() string {
	return "transactions.transactions"
}

// MarshalJSON personaliza la serialización JSON de la entidad Transactions.
// Si el ID es 0, devuelve "null". De lo contrario, serializa la estructura completa.
func (s *Transactions) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Transactions
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
