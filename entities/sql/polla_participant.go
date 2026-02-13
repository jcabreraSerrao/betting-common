package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PollaParticipant struct {
	ID                 uint64           `gorm:"primaryKey;autoIncrement" json:"id"`
	PollaID            uint64           `gorm:"column:polla_id;index;not null" json:"pollaId"`
	Polla              Polla            `gorm:"foreignKey:PollaID;references:ID" json:"polla"`
	TercioID           uint64           `gorm:"column:tercio_id;index;not null" json:"tercioId"`
	Tercio             Tercios          `gorm:"foreignKey:TercioID;references:ID" json:"tercio"`
	Ticket             string           `gorm:"column:ticket;unique;not null" json:"ticket"`
	TotalPoints        int              `gorm:"column:total_points;default:0" json:"totalPoints"`
	Prize              decimal.Decimal  `gorm:"column:prize;type:decimal(18,8);default:0" json:"prize"`
	EntryTransactionID uint64           `gorm:"column:entry_transaction_id" json:"entryTransactionId"`
	EntryTransaction   Transactions     `gorm:"foreignKey:EntryTransactionID;references:ID" json:"entryTransaction"`
	PrizeTransactionID *uint64          `gorm:"column:prize_transaction_id" json:"prizeTransactionId"`
	PrizeTransaction   *Transactions    `gorm:"foreignKey:PrizeTransactionID;references:ID" json:"prizeTransaction"`
	Selections         []PollaSelection `gorm:"foreignKey:PollaParticipantID;references:ID" json:"selections"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PollaParticipant) TableName() string {
	return "gaming.polla_participant"
}
