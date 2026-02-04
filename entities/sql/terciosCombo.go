package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TerciosCombo struct {
	ID            uint            `gorm:"primarykey" json:"id"`
	IDCombo       uint            `gorm:"column:id_combo" json:"idCombo"`
	Combo         ComboRemate     `gorm:"foreignKey:IDCombo;references:ID" json:"combo"`
	IDTercios     uint            `gorm:"column:id_tercios" json:"idTercios"`
	Tercios       Tercios         `gorm:"foreignKey:IDTercios;references:ID" json:"tercios"`
	IDTransaction uint64          `gorm:"column:id_transaction" json:"idTransaction"`
	Transaction   Transactions    `gorm:"foreignKey:IDTransaction;references:ID" json:"transaction"`
	Amount        decimal.Decimal `gorm:"column:amount;type:decimal(18,8);default:0" json:"amount"`
	Reverse       bool            `gorm:"column:reverse;default:false;not null" json:"reverse"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (s *TerciosCombo) TableName() string {
	return "gaming.tercios_combo"
}
func (s *TerciosCombo) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TerciosCombo
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
func (s *TerciosCombo) UnmarshalJSON(data []byte) error {
	type Alias TerciosCombo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// If ID is 0, set it to nil
	if s.ID == 0 {
		s = nil
		return nil
	}
	// If ID is not 0, set the DeletedAt field to gorm.DeletedAt{}
	if s.ID != 0 {
		s.DeletedAt = gorm.DeletedAt{}
	}
	return nil
}
