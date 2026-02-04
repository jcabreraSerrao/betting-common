package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

// BancaGroupLink representa la relación entre una banca y un grupo no banca
type BancaGroupLink struct {
	gorm.Model
	BancaID uint `gorm:"column:banca_id;not null;index" json:"bancaId"` // ID del grupo que es banca
	GroupID uint `gorm:"column:group_id;not null;index" json:"groupId"` // ID del grupo no banca

	// Relaciones
	Banca Group `gorm:"foreignKey:BancaID;references:ID" json:"banca"`
	Group Group `gorm:"foreignKey:GroupID;references:ID" json:"group"`

	// Flag para completar deuda
	CompletarDeuda bool `gorm:"column:completar_deuda;not null;default:false" json:"completarDeuda"`
	// Indica si el pago es por porcentaje
	EsPorcentaje bool `gorm:"column:es_porcentaje;not null;default:false" json:"esPorcentaje"`
	// Monto del porcentaje (si aplica)
	MontoPorcentaje *float64 `gorm:"column:monto_porcentaje" json:"montoPorcentaje,omitempty"`
	// Monto fijo (si aplica)
	MontoFijo *float64 `gorm:"column:monto_fijo" json:"montoFijo,omitempty"`
}

func (b *BancaGroupLink) TableName() string {
	return "security.banca_group_link"
}

func (b *BancaGroupLink) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}
	type Alias BancaGroupLink
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
		},
	)
}
