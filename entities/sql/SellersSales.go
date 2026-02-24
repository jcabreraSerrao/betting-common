package sql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// SellersSales representa la vista reports.sellers_sales
// Contiene información de ventas/amounts para reportes de vendedores
type SellersSales struct {
	IDGroup         uint            `gorm:"column:id_group" json:"id_group"`
	Date            time.Time       `gorm:"column:date" json:"date"`
	CurrencyName    string          `gorm:"column:currency_name" json:"currency_name"`
	UserName        *string         `gorm:"column:user_name" json:"user_name"`
	IDTercio        *uint           `gorm:"column:id_tercio" json:"id_tercio"`
	MontoTotalApues decimal.Decimal `gorm:"column:monto_total_apues;type:numeric" json:"monto_total_apues"`
}

// TableName devuelve el nombre de la vista en la base de datos
func (s *SellersSales) TableName() string {
	return "reports.sellers_sales"
}

// BeforeCreate previene la inserción en la vista
func (s *SellersSales) BeforeCreate(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}

// BeforeUpdate previene la actualización en la vista
func (s *SellersSales) BeforeUpdate(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}

// BeforeDelete previene la eliminación en la vista
func (s *SellersSales) BeforeDelete(tx *gorm.DB) error {
	return gorm.ErrInvalidDB // Las vistas son solo lectura
}
