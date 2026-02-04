package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ComboRemate struct {
	ID               uint               `gorm:"primarykey" json:"id"`
	IDRemate         uint               `gorm:"column:id_remate" json:"idRemate"`
	Remate           ConfigRemate       `gorm:"foreignKey:IDRemate;references:ID" json:"remate"`
	Invalidated      bool               `gorm:"column:invalidated;default:false" json:"invalidated"`
	ParticipantCombo []ParticipantCombo `gorm:"foreignKey:ComboID;references:ID" json:"participantCombo"`
	TerciosCombo     []TerciosCombo     `gorm:"foreignKey:IDCombo;references:ID" json:"terciosCombo"`
	CreatedAt        time.Time          `json:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"-"`
}

func (s *ComboRemate) TableName() string {
	return "gaming.combo_remate"
}
func (s *ComboRemate) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias ComboRemate
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
func (s *ComboRemate) UnmarshalJSON(data []byte) error {
	type Alias ComboRemate
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// If ID is 0, set it to nil
	if s.ID == 0 {
		s = nil
		return nil
	}
	// If ID is not 0, set the DeletedAt field to gorm.DeletedAt{}
	if s.ID != 0 {
		s.DeletedAt = gorm.DeletedAt{}
	}
	return nil
}
