package sql

import "github.com/shopspring/decimal"

/* dia
id_tercios
name
id_tercio
total
id_group */

type ViewSaldoDayTercio struct {
	Day      string          `gorm:"column:dia" json:"day"`
	IDTercio int             `gorm:"column:id_tercios" json:"idTercio"`
	Name     string          `gorm:"column:name" json:"name"`
	IDGroup  int             `gorm:"column:id_group" json:"idGroup"`
	Total    decimal.Decimal `gorm:"column:total" json:"total"`
}

// TableName especifica el nombre de la vista en la base de datos.
func (ViewSaldoDayTercio) TableName() string {
	return "reports.tercios_diarios"
}
