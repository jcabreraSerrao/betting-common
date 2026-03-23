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
	ErrorInsufficientBalance   ErrorWarningType = "INSUFFICIENT_BALANCE"
	ErrorTercioNotFound        ErrorWarningType = "TERCIO_NOT_FOUND"
	ErrorTercioMakerNotFound   ErrorWarningType = "TERCIO_MAKER_NOT_FOUND"
	ErrorTercioTakerNotFound   ErrorWarningType = "TERCIO_TAKER_NOT_FOUND"
	ErrorInvalidFormat         ErrorWarningType = "INVALID_FORMAT"
	ErrorNLUFailure            ErrorWarningType = "NLU_FAILURE"
	ErrorAutomaticBetFailure   ErrorWarningType = "AUTOMATIC_BET_FAILURE"
	ErrorRaceClosed            ErrorWarningType = "RACE_CLOSED"
	ErrorUnexpected            ErrorWarningType = "UNEXPECTED_ERROR"
	ErrorMatchInProgress       ErrorWarningType = "MATCH_IN_PROGRESS"
	ErrorAlreadyProcessed      ErrorWarningType = "ALREADY_PROCESSED"
)

// MatchAttemptResponse representa una respuesta individual (intento) a un mensaje maestro.
type MatchAttemptResponse struct {
	MessageID       string           `json:"message_id"`
	ParticipantJID  string           `json:"participant_jid"` // LID Limpio (ej: 12345@lid)
	Command         string           `json:"command"`         // Lo que el usuario escribió
	Status          string           `json:"status"`          // SUCCESS, ERROR
	ErrorType       ErrorWarningType `json:"error_type,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
	ResponseMessage string           `json:"response_message,omitempty"` // Lo que respondió el bot
	Timestamp       time.Time        `json:"timestamp"`
}

// WhatsappMatchAttempt agrupa todos los intentos de respuesta bajo un mismo mensaje original (Maker).
type WhatsappMatchAttempt struct {
	IDLogs            uuid.UUID      `gorm:"primaryKey;column:id_logs;type:uuid;default:uuid_generate_v4()" json:"id_logs"`
	RaceID            uint64         `gorm:"column:race_id;index:idx_wa_attempt_race" json:"race_id"`
	GroupID           uint           `gorm:"column:group_id;index:idx_wa_attempt_group_id" json:"group_id"`
	QuotedMessageID   string         `gorm:"column:quoted_message_id;type:varchar(255);uniqueIndex:idx_wa_attempt_quoted_id" json:"quoted_message_id"`
	QuotedParticipant string         `gorm:"column:quoted_participant;type:varchar(255)" json:"quoted_participant"`
	QuotedBody        string         `gorm:"column:quoted_body;type:text" json:"quoted_body"`
	Responses         datatypes.JSON `gorm:"column:responses" json:"responses"` // Arreglo de MatchAttemptResponse
	CreatedAt         time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (WhatsappMatchAttempt) TableName() string {
	return "whatsapp.match_attempts"
}
