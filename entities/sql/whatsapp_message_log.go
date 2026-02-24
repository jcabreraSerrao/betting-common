package sql

import (
	"time"

	"gorm.io/gorm"
)

// WhatsappMessageLog almacena los logs específicos de juego/comandos para auditoría a largo plazo.
// Separado de los mensajes generales (ProcessedMessage) que pueden ser purgados.
type WhatsappMessageLog struct {
	MessageID         string         `gorm:"primaryKey;column:message_id;type:varchar(255)" json:"message_id"`
	RemoteJid         string         `gorm:"column:remote_jid;type:varchar(255);index:idx_evo_msg_remote_jid" json:"remote_jid"`
	Participant       string         `gorm:"column:participant;type:varchar(255)" json:"participant"`
	Command           string         `gorm:"column:command;type:varchar(255)" json:"command"`
	Params            string         `gorm:"column:params;type:text" json:"params"`
	QuotedMessageID   string         `gorm:"column:quoted_message_id;type:varchar(255)" json:"quoted_message_id"`
	QuotedParticipant string         `gorm:"column:quoted_participant;type:varchar(255)" json:"quoted_participant"`
	QuotedBody        string         `gorm:"column:quoted_body;type:text" json:"quoted_body"`
	Tercio1ID         uint           `gorm:"column:tercio_1_id;index" json:"tercio_1_id"`
	Tercio2ID         uint           `gorm:"column:tercio_2_id;index" json:"tercio_2_id"`
	Timestamp         time.Time      `gorm:"column:timestamp;index:idx_evo_msg_timestamp" json:"timestamp"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (WhatsappMessageLog) TableName() string {
	return "whatsapp.message_logs"
}
