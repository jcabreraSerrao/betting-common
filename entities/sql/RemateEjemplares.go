package sql

import (
	"encoding/json"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RemateEjemplares struct {
	Ejemplar string          `gorm:"column:ejemplar" json:"ejemplar"`
	Amount   decimal.Decimal `gorm:"column:amount;type:decimal(18,8)" json:"amount"`
	Alias    string          `gorm:"column:alias" json:"alias"`
	RemateID uint            `gorm:"column:id_remate" json:"idRemate"`
	Remate   ConfigRemate    `gorm:"foreignKey:RemateID;references:ID" json:"remate"`
	gorm.Model
}

func (*RemateEjemplares) TableName() string {
	return "gaming.remate_ejemplares"
}

func (r *RemateEjemplares) MarshalJSON() ([]byte, error) {
	type Alias RemateEjemplares
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

func (r *RemateEjemplares) UnmarshalJSON(data []byte) error {
	type Alias RemateEjemplares
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
