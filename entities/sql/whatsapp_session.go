package sql

import (
	"time"

	"gorm.io/gorm"
)

// WhatsAppSession representa una sesión de whatsmeow
type WhatsAppSession struct {
	ID              uint           `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
	SessionID       string         `gorm:"column:session_id;type:varchar(100);uniqueIndex;not null" json:"session_id"`
	GroupID         uint           `gorm:"column:group_id;index;not null" json:"group_id"` // ID del grupo en nuestra BD
	Group           *Group         `gorm:"foreignKey:GroupID"`
	ProxyID         *uint          `gorm:"column:proxy_id;index" json:"proxy_id,omitempty"` // Proxy asignado
	Proxy           *WhatsAppProxy `gorm:"foreignKey:ProxyID"`
	JID             string         `gorm:"column:jid;type:varchar(100)" json:"jid,omitempty"`              // JID del usuario conectado
	Status          string         `gorm:"column:status;type:varchar(20);default:'pending'" json:"status"` // pending, connected, disconnected
	RegistrationID  *uint32        `gorm:"column:registration_id" json:"registration_id,omitempty"`        // ID de registro en whatsmeow
	LastConnectedAt *time.Time     `gorm:"column:last_connected_at" json:"last_connected_at,omitempty"`
	QRExpiresAt     *time.Time     `gorm:"column:qr_expires_at" json:"qr_expires_at,omitempty"`
}

func (WhatsAppSession) TableName() string {
	return "config.whatsapp_session"
}
