package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Refills struct {
	gorm.Model
	Transaction            Transactions          `gorm:"foreignKey:TransactionId;references:ID"`
	TransactionId          uint                  `gorm:"column:id_transaction"`
	PaymentPlatformGroup   PaymentPlatformsGroup `gorm:"foreignKey:PaymentPlatformGroupId;references:ID"`
	PaymentPlatformGroupId uint                  `gorm:"column:id_payment_platform_group"`
	Approved               bool                  `gorm:"column:approved;type:bool;default:false;not null"`
	Code                   string                `gorm:"column:code;type:varchar(255);not null"`
	Datos                  json.RawMessage       `gorm:"column:datos;type:json" json:"datos"`
	Group                  Group                 `gorm:"foreignKey:GroupId;references:ID"`
	GroupId                uint                  `gorm:"column:id_group"`
	User                   UserGroup             `gorm:"foreignKey:UserId;references:ID"`
	UserId                 uint                  `gorm:"column:id_user"`
}

func (r *Refills) TableName() string {
	return "transactions.refills"
}

func (r *Refills) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Refills
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
