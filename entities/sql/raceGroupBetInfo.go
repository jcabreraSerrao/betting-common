package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// RaceGroupCommission representa las comisiones y conteo de apuestas por carrera y grupo.
type RaceGroupCommission struct {
	ID              uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	RaceID          uint64          `gorm:"column:id_race;not null" json:"idRace"`
	Race            Race            `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	GroupID         uint            `gorm:"column:id_group;not null" json:"idGroup"`
	Group           Group           `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	TotalCommission decimal.Decimal `gorm:"column:total_commission;type:decimal(18,8);default:0" json:"totalCommission"`
	BetCount        int             `gorm:"column:bet_count;default:0" json:"betCount"`
	TotalAmount     decimal.Decimal `gorm:"column:total_amount;type:decimal(18,8);default:0" json:"totalAmount"`
	RaceDate        *time.Time      `gorm:"column:race_date" json:"raceDate"`
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt  `gorm:"index;column:deleted_at" json:"-"`
}

func (r *RaceGroupCommission) TableName() string {
	return "gaming.race_group_bet_info"
}

func (r *RaceGroupCommission) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceGroupCommission
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

func (r *RaceGroupCommission) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	type Alias RaceGroupCommission
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(data, aux)
}
