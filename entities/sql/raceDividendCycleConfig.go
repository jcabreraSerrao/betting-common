package sql

import (
	"time"

	"gorm.io/gorm"
)

// RaceDividendCycleConfig: configuración por ciclo y rango de participantes para RaceDividendConfig
// Relaciona con RaceDividendConfig por ConfigID
// gaming.race_dividend_cycle_config

type RaceDividendCycleConfig struct {
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Ciclo              string         `gorm:"column:ciclo;not null" json:"ciclo"` // string para identificar el ciclo
	ParticipantesDesde int            `gorm:"column:participantes_desde;not null" json:"participantesDesde"`
	ParticipantesHasta int            `gorm:"column:participantes_hasta;not null" json:"participantesHasta"`
	MontoExtraEmpate   float64        `gorm:"column:monto_extra_empate;type:decimal(14,2)" json:"montoExtraEmpate"`
	TopeDividendo      float64        `gorm:"column:tope_dividendo;type:decimal(14,2)" json:"topeDividendo"`
	Status             bool           `gorm:"column:status;default:true" json:"status"`
	GroupID            *uint          `gorm:"column:group_id" json:"groupId"`                                                                        // Optional association to group
	Group              *Group         `gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group,omitempty"` // Associated group
	CreatedAt          time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt          gorm.DeletedAt `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

func (r *RaceDividendCycleConfig) TableName() string {
	return "gaming.race_dividend_cycle_config"
}
