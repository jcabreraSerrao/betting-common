package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// SellersGeneralReport representa la vista reports.sellers_general_report
// Contiene información de premios, reembolsos y pendientes para reportes de vendedores
type SellersGeneralReport struct {
	IDGroup             uint            `gorm:"column:id_group" json:"id_group"`
	Date                time.Time       `gorm:"column:date" json:"date"`
	CurrencyName        string          `gorm:"column:currency_name" json:"currency_name"`
	UserName            *string         `gorm:"column:user_name" json:"user_name"`
	IDTercio            *uint           `gorm:"column:id_tercio" json:"id_tercio"`
	MontoPrizes         decimal.Decimal `gorm:"column:monto_prizes;type:numeric" json:"monto_prizes"`
	MontoRefunds        decimal.Decimal `gorm:"column:monto_refunds;type:numeric" json:"monto_refunds"`
	MontoPending        decimal.Decimal `gorm:"column:monto_pending;type:numeric" json:"monto_pending"`
	MontoPendingRefunds decimal.Decimal `gorm:"column:monto_pending_refunds;type:numeric" json:"monto_pending_refunds"`
}

// TableName devuelve el nombre de la vista en la base de datos
func (s *SellersGeneralReport) TableName() string {
	return "reports.sellers_general_report"
}

// BeforeCreate previene la inserción en la vista
func (s *SellersGeneralReport) BeforeCreate(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}

// BeforeUpdate previene la actualización en la vista
func (s *SellersGeneralReport) BeforeUpdate(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}

// BeforeDelete previene la eliminación en la vista
func (s *SellersGeneralReport) BeforeDelete(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}
