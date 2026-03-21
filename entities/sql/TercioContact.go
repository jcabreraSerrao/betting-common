package sql

import (
	"time"

	"gorm.io/gorm"
)

type ContactType string

const (
	ContactTypePhone    ContactType = "PHONE"
	ContactTypeWhatsApp ContactType = "WHATSAPP"
	ContactTypeTelegram ContactType = "TELEGRAM"
)

type TercioContact struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	TercioID     uint64         `gorm:"column:id_tercio;index" json:"id_tercio"`
	ContactType  ContactType    `gorm:"column:contact_type" json:"contact_type"`
	ContactValue string         `gorm:"column:contact_value;index" json:"contact_value"`
	IsPrimary    bool           `gorm:"column:is_primary;default:false" json:"is_primary"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (TercioContact) TableName() string {
	return "gaming.tercio_contacts"
}
