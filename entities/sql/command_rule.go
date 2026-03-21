package sql

import (
	"gorm.io/gorm"
)

// CommandRule representa una regla de detección de comandos dinámicos.
// Pueden ser globales (GroupID == nil) o específicas de un grupo.
type CommandRule struct {
	gorm.Model
	CommandType   string   `gorm:"column:command_type;not null;index:idx_command_rules_type_group" json:"command_type"`
	Aliases       JSONBMap `gorm:"column:aliases;type:jsonb;not null" json:"aliases"`
	MinSimilarity float64  `gorm:"column:min_similarity;default:0.8" json:"min_similarity"`
	GroupID       *uint    `gorm:"column:group_id;index:idx_command_rules_type_group" json:"group_id"` // Nullable: Global rule
	IsActive      bool     `gorm:"column:is_active;default:true" json:"is_active"`
}

// TableName define el nombre de la tabla en el esquema whatsapp
func (CommandRule) TableName() string {
	return "whatsapp.command_rules"
}
