package sql

import "time"

// WorkingDay representa una jornada de operación de un grupo.
// Controla el período durante el cual se permiten apuestas, recargas, retiros y creación de tercios.
// Los saldos de los tercios se calculan usando open_date/close_date como rango de timestamps exactos.
type WorkingDay struct {
	ID        uint64               `gorm:"primaryKey;autoIncrement;column:id"             json:"id"`
	GroupID   uint                 `gorm:"not null;index;column:id_group"                 json:"id_group"`
	Name      string               `gorm:"not null;default:'';column:name"                json:"name"`
	OpenDate  time.Time            `gorm:"not null;column:open_date;type:timestamptz"     json:"open_date"`
	CloseDate *time.Time           `gorm:"default:null;column:close_date;type:timestamptz" json:"close_date"`
	Open      bool                 `gorm:"not null;default:true;column:open"              json:"open"`
	Snapshots []WorkingDaySnapshot `gorm:"foreignKey:WorkingDayID"                        json:"snapshots,omitempty"`
	CreatedAt time.Time            `gorm:"column:created_at;type:timestamptz;default:now()" json:"created_at"`
	UpdatedAt time.Time            `gorm:"column:updated_at;type:timestamptz;default:now()" json:"updated_at"`
}

func (WorkingDay) TableName() string {
	return "gaming.working_days"
}
