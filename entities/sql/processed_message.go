package sql

import (
	"time"

	"github.com/google/uuid"
)

// ProcessedMessage representa un mensaje de WhatsApp que ya fue procesado
// Esta tabla usa particionado por rango de fecha (processed_at) para mantener queries rápidos
type ProcessedMessage struct {
	ID               uuid.UUID  `gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	MessageID        string     `gorm:"column:message_id;type:varchar(255);not null" json:"message_id"`                // ID único del mensaje de WhatsApp
	ChatJID          string     `gorm:"column:chat_jid;type:varchar(255);not null;index:idx_chat_jid" json:"chat_jid"` // JID del chat (grupo)
	SessionID        string     `gorm:"column:session_id;type:varchar(100);not null" json:"session_id"`                // Sesión que procesó el mensaje
	GroupID          uint       `gorm:"column:group_id;not null;index:idx_group_id" json:"group_id"`                   // ID interno del grupo
	SenderJID        string     `gorm:"column:sender_jid;type:varchar(255);not null" json:"sender_jid"`                // JID del remitente
	TextContent      string     `gorm:"column:text_content;type:text" json:"text_content"`                             // Contenido del mensaje (para auditoría)
	ProcessedAt      time.Time  `gorm:"column:processed_at;not null;index:idx_processed_at" json:"processed_at"`       // Timestamp de procesamiento (CLAVE para particionado)
	MessageTimestamp time.Time  `gorm:"column:message_timestamp;not null" json:"message_timestamp"`                    // Timestamp original del mensaje
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (ProcessedMessage) TableName() string {
	return "config.processed_messages"
}
