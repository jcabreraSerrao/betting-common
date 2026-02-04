package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Withdrawal struct {
	ID                     uint                   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt              time.Time              `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt              time.Time              `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt              gorm.DeletedAt         `gorm:"index;column:deleted_at;type:timestamp with time zone"`
	IdTransaction          *uint64                `gorm:"column:id_transaction" json:"id_transaction"`                                                                                          // Renombrado para coincidir con la columna en la DB
	Transaction            *Transactions          `gorm:"foreignKey:IdTransaction;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_transactions_transactions_withdrawals"` // Corregido el nombre de la clave foránea
	IdPaymentPlatformGroup *uint64                `gorm:"column:id_payment_platform_group" json:"id_payment_platform_group"`
	PaymentPlatformGroup   *PaymentPlatformsGroup `gorm:"foreignKey:IdPaymentPlatformGroup;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_transactions_withdrawal_payment_platform_group"`
	Approved               bool                   `gorm:"column:approved;type:boolean;default:false" json:"approved"`
	Code                   string                 `gorm:"column:code;type:varchar(255);not null" json:"code"`
	Datos                  json.RawMessage        `gorm:"column:datos;type:json" json:"datos"` // Cambiado a json.RawMessage para manejar JSON
	IdGroup                *uint64                `gorm:"column:id_group" json:"id_group"`
	Group                  *Group                 `gorm:"foreignKey:IdGroup;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_transactions_withdrawal_group"`
	IdUser                 *uint64                `gorm:"column:id_user" json:"id_user"`
	User                   *UserGroup             `gorm:"foreignKey:IdUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_transactions_withdrawal_user"`
}

func (w *Withdrawal) TableName() string {
	return "transactions.withdrawal"
}

func (w *Withdrawal) MarshalJSON() ([]byte, error) {
	if w.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Withdrawal
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(w),
		},
	)
}
