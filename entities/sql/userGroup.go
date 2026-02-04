package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type UserGroup struct {
	Name          string `gorm:"column:name" json:"name"`
	User          string `gorm:"column:username" json:"user"`
	Password      string `gorm:"column:password" json:"password"`
	Email         string `gorm:"column:email" json:"email"`
	Tlf           string `gorm:"column:phone_number" json:"tlf"`
	RecoveryToken string `gorm:"column:recovery_token" json:"recoveryToken"`
	Status        bool   `gorm:"column:status;default:true" json:"status"`
	IDGroup       uint   `gorm:"column:id_group" json:"idGroup"`
	IDRole        uint   `gorm:"column:id_rol" json:"idRole"`
	TercioID      *uint  `gorm:"column:id_tercio" json:"tercioId"`
	Secret        string `gorm:"column:secret" json:"secret"`
	Rol           Roles  `gorm:"foreignKey:IDRole" json:"rol"`
	Group         Group  `gorm:"foreignKey:IDGroup" json:"group"`
	gorm.Model
}

func (u *UserGroup) TableName() string {
	return "security.user_group" // Se indica el nombre de la tabla en la BD
}

func (u *UserGroup) MarshalJSON() ([]byte, error) {
	if u.ID == 0 {
		return []byte("null"), nil
	}

	type Alias UserGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			// Utiliza el nombre personalizado del ID
			Alias: (*Alias)(u),
		},
	)
}
