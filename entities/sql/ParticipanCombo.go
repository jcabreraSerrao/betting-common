package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ParticipantCombo struct {
	ID             uint             `gorm:"primarykey" json:"id"`
	ComboID        uint             `gorm:"column:id_combo" json:"idCombo"`
	Combo          ComboRemate      `gorm:"foreignKey:ComboID;references:ID" json:"combo"`
	IDParticipanst uint             `gorm:"column:id_participants" json:"idParticipants"`
	Participants   ParticipantsRace `gorm:"foreignKey:IDParticipanst;references:ID"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (s *ParticipantCombo) TableName() string {
	return "gaming.participant_combo"
}

func (s *ParticipantCombo) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias ParticipantCombo
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}

func (s *ParticipantCombo) UnmarshalJSON(data []byte) error {
	type Alias ParticipantCombo
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
