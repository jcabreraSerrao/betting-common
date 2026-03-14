package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ErrorWarningType es un enum para los tipos de errores o advertencias en el procesamiento de mensajes.
type ErrorWarningType string

const (
	ErrorInsufficientBalance  ErrorWarningType = "INSUFFICIENT_BALANCE"
	ErrorTercioNotFound       ErrorWarningType = "TERCIO_NOT_FOUND"
	ErrorInvalidFormat        ErrorWarningType = "INVALID_FORMAT"
	ErrorNLUFailure           ErrorWarningType = "NLU_FAILURE"
	ErrorAutomaticBetFailure  ErrorWarningType = "AUTOMATIC_BET_FAILURE"
	ErrorRaceClosed           ErrorWarningType = "RACE_CLOSED"
	ErrorUnexpected           ErrorWarningType = "UNEXPECTED_ERROR"
	ErrorMatchInProgress      ErrorWarningType = "MATCH_IN_PROGRESS"
	ErrorAlreadyProcessed     ErrorWarningType = "ALREADY_PROCESSED"
)

// WhatsappMatchAttempt registra cada intento individual de procesar un comando o casada.
// Esta tabla permite auditoría detallada, especialmente cuando múltiples personas responden al mismo mensaje.
type WhatsappMatchAttempt struct {
	IDLogs            uuid.UUID        `gorm:"primaryKey;column:id_logs;type:uuid;default:uuid_generate_v4()" json:"id_logs"`
	TraceID           string           `gorm:"column:trace_id;type:varchar(255);index:idx_wa_attempt_trace_id" json:"trace_id"`
	GroupID           uint             `gorm:"column:group_id;index:idx_wa_attempt_group_id" json:"group_id"`
	MessageID         string           `gorm:"column:message_id;type:varchar(255);index:idx_wa_attempt_message_id" json:"message_id"`
	RemoteJid         string           `gorm:"column:remote_jid;type:varchar(255)" json:"remote_jid"`
	Participant       string           `gorm:"column:participant;type:varchar(255)" json:"participant"`
	Command           string           `gorm:"column:command;type:varchar(255)" json:"command"`
	QuotedMessageID   string           `gorm:"column:quoted_message_id;type:varchar(255);index:idx_wa_attempt_quoted_id" json:"quoted_message_id"`
	Status            string           `gorm:"column:status;type:varchar(50)" json:"status"` // SUCCESS, ERROR
	ErrorType         ErrorWarningType `gorm:"column:error_type;type:varchar(100)" json:"error_type"`
	ErrorMessage      string           `gorm:"column:error_message;type:text" json:"error_message"`
	ExternalAPIRequest  datatypes.JSON   `gorm:"column:external_api_request" json:"external_api_request"`
	ExternalAPIResponse datatypes.JSON   `gorm:"column:external_api_response" json:"external_api_response"`
	ResponseMessage   string           `gorm:"column:response_message;type:text" json:"response_message"`
	CreatedAt         time.Time        `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time        `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (WhatsappMatchAttempt) TableName() string {
	return "whatsapp.match_attempts"
}
