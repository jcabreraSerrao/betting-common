package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SettlementRaceEstimate struct {
	ID             uint                           `gorm:"primaryKey;autoIncrement" json:"id"`
	RaceID         uint64                         `gorm:"not null" json:"race_id"`
	GroupID        uint                           `gorm:"not null" json:"group_id"`
	TotalPayout    decimal.Decimal                `gorm:"type:decimal(15,2);not null" json:"total_payout"`
	TotalRefund    decimal.Decimal                `gorm:"type:decimal(15,2);not null" json:"total_refund"`
	TotalFinal     decimal.Decimal                `gorm:"type:decimal(15,2);not null" json:"total_final"`
	TotalCollected decimal.Decimal                `gorm:"type:decimal(15,2);not null;default:0.00" json:"total_collected"`
	CurrencyID     uint                           `gorm:"not null" json:"currency_id"`
	Currency       Currency                       `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`
	Status         string                         `gorm:"type:varchar(20);not null;default:'pending'" json:"status"` // pending, completed, failed
	BetDetails     []SettlementRaceEstimateDetail `gorm:"foreignKey:SettlementRaceEstimateID" json:"bet_details"`    // Relación con detalles de apuestas
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt                 `gorm:"index" json:"-"`
}

func (s *SettlementRaceEstimate) TableName() string {
	return "gaming.settlement_race_estimates"
}
