package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SettlementRaceEstimateDetail struct {
	ID                       uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	SettlementRaceEstimateID uint                   `gorm:"not null" json:"settlement_race_estimate_id"`
	SettlementRaceEstimate   SettlementRaceEstimate `gorm:"foreignKey:SettlementRaceEstimateID;references:ID" json:"settlement_race_estimate"`
	BetID                    uint64                 `gorm:"not null" json:"bet_id"`
	Bet                      Bet                    `gorm:"foreignKey:BetID;references:ID" json:"bet"`
	Payout                   decimal.Decimal        `gorm:"type:decimal(15,2);not null;default:0.00" json:"payout"`
	Refund                   decimal.Decimal        `gorm:"type:decimal(15,2);not null;default:0.00" json:"refund"`
	Amount                   decimal.Decimal        `gorm:"type:decimal(15,2);not null;default:0.00" json:"amount"`
	CreatedAt                time.Time              `json:"created_at"`
	UpdatedAt                time.Time              `json:"updated_at"`
	DeletedAt                gorm.DeletedAt         `gorm:"index" json:"-"`
}

func (s *SettlementRaceEstimateDetail) TableName() string {
	return "gaming.settlement_race_estimate_details"
}
