package sql

import "github.com/shopspring/decimal"

// SaldoTerciosUserGroup mapea la vista reports.saldo_tercios_user_group
type SaldoTerciosUserGroup struct {
	Total      decimal.Decimal `gorm:"column:total" json:"total"`
	IDTercios  uint            `gorm:"column:id_tercios" json:"id_tercios"`
	Name       string          `gorm:"column:name" json:"name"`
	IDGroup    uint            `gorm:"column:id_group" json:"id_group"`
	Status     bool            `gorm:"column:status" json:"status"`
	IDUser     uint            `gorm:"column:id_user" json:"id_user"`
	IDCurrency uint            `gorm:"column:id_currency" json:"id_currency"`
}

// TableName especifica el nombre de la vista en la base de datos.
func (SaldoTerciosUserGroup) TableName() string {
	return "reports.saldo_tercios_user_group"
}
