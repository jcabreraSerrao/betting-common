package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Tercios struct {
	ID           uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string          `gorm:"column:name" json:"name"`
	Slug         string          `gorm:"column:slug" json:"slug"`
	Status       bool            `gorm:"column:status;default:true" json:"status"`
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(18,8)" json:"amount"`
	AmountLock   decimal.Decimal `gorm:"column:amount_lock;type:decimal(18,8)" json:"amount_lock"`
	GroupId      uint            `gorm:"column:id_group" json:"id_group"`
	Group        Group           `gorm:"foreignKey:GroupId;references:ID"`
	TypeTercioID uint            `gorm:"column:id_type_tercio" json:"id_type_tercio"`
	TypeTercio   TypeTercio      `gorm:"foreignKey:TypeTercioID;references:ID"`
	UserTercioID *uint           `gorm:"column:id_user_tercio" json:"id_user_tercio"`
	UserTercio   *UserTercio     `gorm:"foreignKey:UserTercioID;references:ID"`
	Reverso      *TercioReverso  `gorm:"foreignKey:TercioID;references:ID" json:"reverso,omitempty"`
	Transactions []Transactions  `gorm:"foreignKey:TerciosId;"`
	Bet          []Bet           `gorm:"foreignKey:TercioId;"`
	Bet2         []Bet           `gorm:"foreignKey:TercioId2;"`
	TelegramID   string          `gorm:"column:id_telegram" json:"id_telegram"`
	Token        string          `gorm:"column:token" json:"token"`
	Casa         bool            `gorm:"column:casa;default:false" json:"casa"`
	PhoneNumber  string          `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Tercios) TableName() string {
	return "gaming.tercios"
}

func (s *Tercios) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Tercios
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
