package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TerciosRemate struct {
	ID                 uint64            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IdConfigRemate     *uint             `gorm:"column:id_config_remate" json:"idConfigRemate"` // Changed type from *int64 to *uint
	ConfigRemate       *ConfigRemate     `gorm:"foreignKey:IdConfigRemate;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_gaming_config_remate_tercios_remate" json:"configRemate,omitempty"`
	ParticipantsRaceId *uint             `gorm:"column:id_participants_race" json:"idParticipants"`
	ParticipantsRace   *ParticipantsRace `gorm:"foreignKey:ParticipantsRaceId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_gaming_participants_race_tercios_remate" json:"participants"`
	IdTercio           *uint64           `gorm:"column:id_tercio" json:"idTercio"`
	Tercio             *Tercios          `gorm:"foreignKey:IdTercio;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_gaming_tercios_remate_tercio" json:"tercio,omitempty"`
	Amount             decimal.Decimal   `gorm:"column:amount;type:decimal(18,8)" json:"amount"`
	Revert             bool              `gorm:"column:revert;default:false" json:"revert"`
	IdTransaction      uint64            `gorm:"column:id_transaction" json:"idTransaction"`
	Transaction        Transactions      `gorm:"foreignKey:IdTransaction;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"transaction,omitempty"`
	UserId             uint64            `gorm:"column:user_id" json:"userId"` // User who created the remate
	User               UserGroup         `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,fkName:fk_gaming_user_group_tercios_remate" json:"user,omitempty"`
	CreatedAt          time.Time         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updateAt"`
	DeletedAt          gorm.DeletedAt    `gorm:"index;column:deleted_at" json:"deletedAt,omitempty"`
}

func (s *TerciosRemate) TableName() string {
	return "gaming.tercios_remate"
}

func (s *TerciosRemate) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias TerciosRemate
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}

func (s *TerciosRemate) UnmarshalJSON(data []byte) error {
	type Alias TerciosRemate
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
