package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type TypePaymentPlatform struct {
	gorm.Model
	Name             string            `gorm:"column:name;type:varchar(255);not null" json:"name"`
	PaymentPlatforms []PaymentPlatform `gorm:"-" json:"paymentPlatforms"`
}

func (b *TypePaymentPlatform) TableName() string {
	return "payments.type_payment_platform"
}

func (b *TypePaymentPlatform) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TypePaymentPlatform
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
