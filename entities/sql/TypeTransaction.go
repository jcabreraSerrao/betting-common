package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

// TypeTransaction representa un tipo de transacción en el sistema.
// Contiene información sobre el nombre, tipo de cálculo, clasificación y otros detalles del tipo de transacción.
type TypeTransaction struct {
	Name           string         `gorm:"column:name" json:"name"`                     // Nombre del tipo de transacción
	TypeCalc       string         `gorm:"column:type_calc" json:"type_calc"`           // Tipo de cálculo utilizado para este tipo de transacción
	Classification string         `gorm:"column:classification" json:"classification"` // Clasificación del tipo de transacción
	Alias          string         `gorm:"column:alias" json:"alias"`                   // Alias o nombre alternativo
	Transactions   []Transactions `gorm:"foreignKey:TypeTransactionId;references:ID"`  // Transacciones asociadas a este tipo
	Slug           string         `gorm:"column:slug" json:"slug"`                     // Slug para identificación en URLs

	gorm.Model // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
}

// TableName devuelve el nombre de la tabla en la base de datos para la entidad TypeTransaction.
// Este método es utilizado por GORM para mapear la estructura a la tabla correcta.
func (s *TypeTransaction) TableName() string {
	return "transactions.type_transaction"
}

// MarshalJSON personaliza la serialización JSON de la entidad TypeTransaction.
// Si el ID es 0, devuelve "null". De lo contrario, serializa la estructura completa.
func (s *TypeTransaction) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TypeTransaction
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
