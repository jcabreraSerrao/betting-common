package sql

import (
	"time"

	"gorm.io/gorm"
)

// TercioReverso almacena la configuración de reverso de un tercio.
// Soporta tres casos simultáneos:
//   - Caso 1 (self):     SelfPercent → vuelve al mismo tercio
//   - Caso 2 (interno):  InternalTercioID + InternalPercent → otro tercio del mismo grupo
//   - Caso 3 (externo):  ExternalName + ExternalPercent → tercero fuera de la plataforma
type TercioReverso struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// FK al tercio origen (1:1)
	TercioID uint64   `gorm:"column:tercio_id;not null;uniqueIndex" json:"tercio_id"`
	Tercio   *Tercios `gorm:"foreignKey:TercioID;references:ID" json:"-"`

	// Caso 1: self-reverso
	SelfPercent float64 `gorm:"column:self_percent;default:0" json:"self_percent"`

	// Caso 2: tercio interno
	InternalTercioID *uint64  `gorm:"column:internal_tercio_id" json:"internal_tercio_id"`
	InternalTercio   *Tercios `gorm:"foreignKey:InternalTercioID;references:ID" json:"internal_tercio,omitempty"`
	InternalPercent  float64  `gorm:"column:internal_percent;default:0" json:"internal_percent"`

	// Caso 3: tercero externo (solo nombre)
	ExternalName    string  `gorm:"column:external_name" json:"external_name"`
	ExternalPercent float64 `gorm:"column:external_percent;default:0" json:"external_percent"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *TercioReverso) TableName() string {
	return "gaming.tercio_reverso"
}
