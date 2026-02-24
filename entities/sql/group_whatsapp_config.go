package sql

import (
	"time"

	"gorm.io/gorm"
)

type GroupWhatsAppConfig struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID     uint           `gorm:"column:id_group;uniqueIndex" json:"id_group"`
	Group       *Group         `gorm:"foreignKey:GroupID;references:ID"`
	SessionName string         `gorm:"column:session_name" json:"session_name"`
	IsActive    bool           `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (GroupWhatsAppConfig) TableName() string {
	return "config.group_whatsapp_config"
}
