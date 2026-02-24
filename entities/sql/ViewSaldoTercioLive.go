package sql

import (
	"time"

	"github.com/shopspring/decimal"
)

// ViewSaldoTercioLive es una entidad de solo lectura que mapea la vista reports.v_saldo_tercio_live.
//
// REEMPLAZA a ViewSaldoTercio (reports.saldo_tercio).
// El saldo se calcula incrementalmente:
//
//	saldo_actual = ultimo_snapshot.balance + delta(snapshot_at → NOW())
//
// Incluye los mismos campos que ViewSaldoTercio para compatibilidad con repositorios existentes.
// Diferencias de nombres de campo vs ViewSaldoTercio:
//
//	Total       → SaldoActual      (column: total       → saldo_actual)
//	IDTercio    → IDTercio         (column: id_tercios  → id_tercio)
//	IDGroup     → GroupID          (column: id_group    → id_group)
type ViewSaldoTercioLive struct {
	// Saldo calculado (snapshot + delta jornada activa)
	SaldoActual            decimal.Decimal `gorm:"column:saldo_actual"             json:"total"`
	GananciaComissionTotal decimal.Decimal `gorm:"column:ganancia_comission_total" json:"ganancia_comission_total"`
	JornadaOpenDate        time.Time       `gorm:"column:jornada_open_date"        json:"jornada_open_date"`
	UltimoSnapshotAt       *time.Time      `gorm:"column:ultimo_snapshot_at"       json:"ultimo_snapshot_at"`

	// Campos de tercio (mismos que ViewSaldoTercio para compatibilidad)
	IDTercio    uint64 `gorm:"column:id_tercio"    json:"idTercio"`
	Name        string `gorm:"column:name"         json:"name"`
	GroupID     uint   `gorm:"column:id_group"     json:"idGroup"`
	Status      bool   `gorm:"column:status"       json:"status"`
	Token       string `gorm:"column:token"        json:"token"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	IDTelegram  string `gorm:"column:id_telegram"  json:"id_telegram"`
	IDUser      int    `gorm:"column:id_user"      json:"idUser"`
}

func (ViewSaldoTercioLive) TableName() string {
	return "reports.v_saldo_tercio_live"
}
