package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Group struct {
	Name           string      `gorm:"column:name" json:"name"`
	Slug           string      `gorm:"column:slug" json:"slug"`
	IDTelegram     string      `gorm:"column:id_telegram" json:"id_telegram"`
	ChatInstanceID string      `gorm:"column:chat_instance_id" json:"chat_instance_id"`
	Phone          string      `gorm:"column:phone" json:"phone"`
	Status         bool        `gorm:"column:status;default:true" json:"status"`
	Roles          []*Roles    `gorm:"foreignKey:GroupID;references:ID"`
	Tercios        []*Tercios  `gorm:"foreignKey:GroupId;references:ID"`
	SubGroups      []*SubGroup `gorm:"foreignKey:GroupID;references:ID"`
	gorm.Model
	// Indica si el grupo es una banca
	IsBanca bool `gorm:"not null;default:false"`
	// Relaciones con banca y grupos asociados
	BancasAsociadas []BancaGroupLink `gorm:"foreignKey:GroupID;references:ID" json:"bancasAsociadas"`
	GruposAsociados []BancaGroupLink `gorm:"foreignKey:BancaID;references:ID" json:"gruposAsociados"`
	// Código único autogenerado en base de datos
	Code           string               `gorm:"column:code;unique;type:varchar(20)" json:"code"`
	WhatsAppConfig *GroupWhatsAppConfig `gorm:"foreignKey:GroupID"`
}

func (g *Group) TableName() string {
	return "security.group"
}

func (g *Group) MarshalJSON() ([]byte, error) {
	if g.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Group
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(g),
		},
	)
}

/*
	 func (g Group) MarshalJSON() ([]byte, error) {
		type Alias Group
		return json.Marshal(&struct {
			Id uint `json:"id"`
			*Alias
		}{
			Id:    g.ID, // Utiliza el nombre personalizado del ID
			Alias: (*Alias)(&g),
		})
	}
*/
