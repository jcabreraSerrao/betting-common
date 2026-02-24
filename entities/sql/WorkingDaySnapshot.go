package sql

import (
	"time"

	"github.com/shopspring/decimal"
)

// WorkingDaySnapshot guarda el saldo de un tercio al momento exacto de cierre de jornada.
// Funciona como punto de control (checkpoint) para el cálculo incremental de saldos:
//
//	saldo_actual_jornada = Balance (snapshot previo) + delta(transacciones desde open_date)
//
// Nunca se modifica una vez creado; es inmutable por diseño.
type WorkingDaySnapshot struct {
	ID                      uint64          `gorm:"primaryKey;autoIncrement;column:id"                          json:"id"`
	WorkingDayID            uint64          `gorm:"not null;index;column:id_working_day"                        json:"id_working_day"`
	TercioID                uint64          `gorm:"not null;column:id_tercio"                                   json:"id_tercio"`
	GroupID                 uint            `gorm:"not null;index;column:id_group"                              json:"id_group"`
	Balance                 decimal.Decimal `gorm:"type:decimal(18,8);column:balance"                           json:"balance"`                   // saldo neto acumulado al cierre
	GananciaComissionParada decimal.Decimal `gorm:"type:decimal(18,8);column:ganancia_comission_parada"         json:"ganancia_comission_parada"` // suma de commission de la jornada
	SnapshotAt              time.Time       `gorm:"not null;column:snapshot_at;type:timestamptz"               json:"snapshot_at"`                // = CloseDate de la jornada (timestamp exacto)
	CreatedAt               time.Time       `gorm:"column:created_at;type:timestamptz;default:now()"           json:"created_at"`
}

func (WorkingDaySnapshot) TableName() string {
	return "gaming.working_day_snapshots"
}
