package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// RacesProcessGroup representa el proceso de una apuesta en un grupo de carreras, incluyendo su estado y detalles asociados.
type RacesProcessGroup struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`                 // Identificador único del proceso de apuesta
	RaceID    uint64         `gorm:"column:race_id" json:"race_id"`                      // Identificador de la carrera asociada
	Race      Race           `gorm:"foreignKey:RaceID" json:"race"`                      // Carrera asociada
	GroupID   uint64         `gorm:"column:group_id" json:"group_id"`                    // Identificador del grupo asociado
	Group     Group          `gorm:"foreignKey:GroupID" json:"group"`                    // Grupo asociado
	Revert    bool           `gorm:"column:revert;default:false;not null" json:"revert"` // Indica si el proceso ha sido revertido
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`                // Fecha y hora de creación
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`                // Fecha y hora de última actualización
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                     // Fecha y hora de eliminación (para soft delete)
}

// TableName devuelve el nombre de la tabla en la base de datos para la entidad RacesProcessGroup.
func (rpg *RacesProcessGroup) TableName() string {
	return "gaming.races_process_groups"
}

func (rpg *RacesProcessGroup) MarshalJSON() ([]byte, error) {
	type Alias RacesProcessGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(rpg),
		},
	)
}
