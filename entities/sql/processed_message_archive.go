package sql

import (
	"time"

	"github.com/google/uuid"
)

// ProcessedMessageArchive representa mensajes antiguos (>7 días) archivados
// Esta tabla NO usa particionado, es para almacenamiento histórico de largo plazo
type ProcessedMessageArchive struct {
	ID               uuid.UUID `gorm:"column:id;type:uuid;primary_key" json:"id"`
	MessageID        string    `gorm:"column:message_id;type:varchar(255);not null" json:"message_id"`
	ChatJID          string    `gorm:"column:chat_jid;type:varchar(255);not null;index:idx_archive_chat_jid" json:"chat_jid"`
	SessionID        string    `gorm:"column:session_id;type:varchar(100);not null" json:"session_id"`
	GroupID          uint      `gorm:"column:group_id;not null;index:idx_archive_group_id" json:"group_id"`
	SenderJID        string    `gorm:"column:sender_jid;type:varchar(255);not null" json:"sender_jid"`
	TextContent      string    `gorm:"column:text_content;type:text" json:"text_content"`
	ProcessedAt      time.Time `gorm:"column:processed_at;not null;index:idx_archive_processed_at" json:"processed_at"`
	MessageTimestamp time.Time `gorm:"column:message_timestamp;not null" json:"message_timestamp"`
	ArchivedAt       time.Time `gorm:"column:archived_at;not null;default:now()" json:"archived_at"` // Timestamp de archivado
}

func (ProcessedMessageArchive) TableName() string {
	return "config.processed_messages_archive"
}
