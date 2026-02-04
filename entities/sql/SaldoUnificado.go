package sql

import "github.com/shopspring/decimal"

type ViewSaldoTercio struct {
	Total       decimal.Decimal `gorm:"column:total" json:"total"`
	IDTercio    int             `gorm:"column:id_tercios" json:"idTercio"`
	Name        string          `gorm:"column:name" json:"name"`
	IDGroup     int             `gorm:"column:id_group" json:"idGroup"`
	Status      bool            `gorm:"column:status" json:"status"`
	Token       string          `gorm:"column:token" json:"token"`
	PhoneNumber string          `gorm:"column:phone_number" json:"phone_number"`
	IDTelegram  string          `gorm:"column:id_telegram" json:"id_telegram"`
	IDUser      int             `gorm:"column:id_user" json:"idUser"`
}

// TableName especifica el nombre de la vista en la base de datos.
func (ViewSaldoTercio) TableName() string {
	return "reports.saldo_tercio"
}
