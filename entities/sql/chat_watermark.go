package sql

import (
	"time"

	"github.com/google/uuid"
)

// ChatWatermark trackea el último mensaje procesado por chat
// Permite recuperación inteligente después de una caída del sistema
type ChatWatermark struct {
	ID                     uuid.UUID  `gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ChatJID                string     `gorm:"column:chat_jid;type:varchar(255);uniqueIndex;not null" json:"chat_jid"`                       // JID del chat
	GroupID                uint       `gorm:"column:group_id;not null;index:idx_watermark_group_id" json:"group_id"`                        // ID interno del grupo
	LastProcessedMessageID string     `gorm:"column:last_processed_message_id;type:varchar(255);not null" json:"last_processed_message_id"` // Último MessageID procesado
	LastProcessedAt        time.Time  `gorm:"column:last_processed_at;not null" json:"last_processed_at"`                                   // Timestamp del último procesamiento
	MessageCount           uint64     `gorm:"column:message_count;default:0" json:"message_count"`                                          // Contador de mensajes procesados
	CreatedAt              time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt              *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (ChatWatermark) TableName() string {
	return "config.chat_watermarks"
}
