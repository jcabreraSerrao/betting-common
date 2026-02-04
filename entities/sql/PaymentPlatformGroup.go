package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type PaymentPlatformsGroup struct {
	gorm.Model
	Name               string          `gorm:"column:name;type:varchar(255);not null" json:"name"`
	PaymentPlatformsID int             `gorm:"column:payment_platform_id;not null" json:"payment_platform_id" binding:"required"`
	PaymentPlatforms   PaymentPlatform `gorm:"foreignKey:payment_platform_id;references:ID;column:payment_platform_id" json:"payment_platform"`
	Datos              json.RawMessage `gorm:"column:datos;type:json" json:"datos"`
	Status             bool            `gorm:"column:status;default:true;not null" json:"status"`
	Group              Group           `gorm:"foreignKey:GroupId;references:ID" json:"group"`
	GroupId            int             `gorm:"column:id_group" json:"idGroup"`
	Refills            []*Refills      `gorm:"foreignKey:PaymentPlatformGroupId;references:ID" json:"refills"`
}

func (b *PaymentPlatformsGroup) TableName() string {
	return "payments.PaymentPlatformGroup"
}

func (b *PaymentPlatformsGroup) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias PaymentPlatformsGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
