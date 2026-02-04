package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type PaymentPlatform struct {
	gorm.Model
	CodeIbp        string              `gorm:"column:code_ibp;type:varchar(5)" json:"code_ibp"`
	Name           string              `gorm:"column:name;type:varchar(255);not null" json:"name"`
	TypePlatform   TypePaymentPlatform `gorm:"foreignKey:TypePlatformID;references:ID" json:"typePlatform"`
	TypePlatformID uint                `gorm:"column:id_type_payment_platform" json:"typePlatformID"`
}

func (b *PaymentPlatform) TableName() string {
	return "payments.PaymentPlatform"
}

func (b *PaymentPlatform) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias PaymentPlatform
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
