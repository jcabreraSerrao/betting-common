package sql

import (
	"time"

	"gorm.io/gorm"
)

// WhatsAppProxy representa los proxies disponibles para las sesiones
type WhatsAppProxy struct {
	ID              uint              `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt       time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
	GroupID         uint              `gorm:"column:group_id;index;not null" json:"group_id"` // Proxy pertenece a un grupo
	Group           *Group            `gorm:"foreignKey:GroupID"`
	ProxyURL        string            `gorm:"column:proxy_url;type:varchar(500);not null" json:"proxy_url"`
	Protocol        string            `gorm:"column:protocol;type:varchar(20)" json:"protocol"` // http, socks5
	IsActive        bool              `gorm:"column:is_active;default:true" json:"is_active"`
	ActiveSessionID *string           `gorm:"column:active_session_id;type:varchar(100);uniqueIndex" json:"active_session_id"` // Sesión usándolo
	Sessions        []WhatsAppSession `gorm:"foreignKey:ProxyID"`
}

func (WhatsAppProxy) TableName() string {
	return "config.whatsapp_proxy"
}
