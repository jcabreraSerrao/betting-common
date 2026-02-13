package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Polla struct {
	ID                uint64             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string             `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description       string             `gorm:"column:description;type:text" json:"description"`
	GroupID           uint               `gorm:"column:group_id;index;not null" json:"groupId"`
	Group             Group              `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	EntryFee          decimal.Decimal    `gorm:"column:entry_fee;type:decimal(18,8);not null" json:"entryFee"`
	CurrencyID        uint               `gorm:"column:currency_id;not null" json:"currencyId"` // Siempre moneda principal
	Currency          Currency           `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`
	Status            string             `gorm:"column:status;default:'open';index" json:"status"` // open, locked, closed, settled, cancelled
	NumRaces          int                `gorm:"column:num_races;default:6" json:"numRaces"`
	StartDate         time.Time          `gorm:"column:start_date" json:"startDate"`
	EndDate           time.Time          `gorm:"column:end_date" json:"endDate"`
	FirstRaceClosed   bool               `gorm:"column:first_race_closed;default:false" json:"firstRaceClosed"`
	TotalPool         decimal.Decimal    `gorm:"column:total_pool;type:decimal(18,8);default:0" json:"totalPool"`
	CommissionAmount  decimal.Decimal    `gorm:"column:commission_amount;type:decimal(18,8);default:0" json:"commissionAmount"`
	NetPool           decimal.Decimal    `gorm:"column:net_pool;type:decimal(18,8);default:0" json:"netPool"`
	CommissionPercent decimal.Decimal    `gorm:"column:commission_percent;type:decimal(5,2)" json:"commissionPercent"`
	ClosedAt          *time.Time         `gorm:"column:closed_at" json:"closedAt"`
	Races             []PollaRace        `gorm:"foreignKey:PollaID;references:ID" json:"races"`
	Participants      []PollaParticipant `gorm:"foreignKey:PollaID;references:ID" json:"participants"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Polla) TableName() string {
	return "gaming.polla"
}
