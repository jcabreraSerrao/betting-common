package sql

import (
	"time"

	"gorm.io/gorm"
)

// BetParticipants representa la relación entre una apuesta y los participantes (caballos) en ella.
// Permite asociar múltiples participantes a una sola apuesta, con un indicador de si es el principal.
type BetParticipants struct {
	ID            uint64           `gorm:"primaryKey;autoIncrement" json:"id"`                        // Identificador único
	BetID         uint64           `gorm:"column:id_bet;not null" json:"betId"`                       // ID de la apuesta
	Bet           Bet              `gorm:"foreignKey:BetID;references:ID" json:"bet"`                 // Relación con la apuesta
	ParticipantID uint             `gorm:"column:id_participant;not null" json:"participantId"`       // ID del participante (caballo)
	Participant   ParticipantsRace `gorm:"foreignKey:ParticipantID;references:ID" json:"participant"` // Relación con el participante
	IsMain        bool             `gorm:"column:is_main;default:false;not null" json:"isMain"`       // Indica si es el participante principal
	CreatedAt     time.Time        // Fecha y hora de creación
	UpdatedAt     time.Time        // Fecha y hora de última actualización
	DeletedAt     gorm.DeletedAt   `gorm:"index" json:"-"` // Fecha y hora de eliminación (para soft delete)
}

// TableName devuelve el nombre de la tabla en la base de datos para la entidad BetParticipants.
func (bp *BetParticipants) TableName() string {
	return "gaming.bet_participants"
}
