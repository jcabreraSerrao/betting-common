package sql

import (
	"time"

	"gorm.io/gorm"
)

type TypeTercio struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name;unique;not null" json:"name"`
	Slug      string         `gorm:"column:slug;unique;not null" json:"slug"`
	Status    bool           `gorm:"column:status;default:true" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *TypeTercio) TableName() string {
	return "gaming.type_tercio"
}
