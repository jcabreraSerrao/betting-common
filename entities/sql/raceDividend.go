package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// RaceDividendConfig configures dividend rules and caps by hippodrome and cycle.
// Can be optionally associated to a group to allow specific configurations per client.
type RaceDividendConfig struct {
	ID           uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	HippodromeID uint        `gorm:"column:hippodrome_id;not null" json:"hippodrome_id"`
	Hippodrome   Hippodromes `gorm:"foreignKey:HippodromeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"hippodrome"`
	GroupID      *uint       `gorm:"column:group_id" json:"group_id"`                                                                       // Optional association to group
	Group        *Group      `gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group,omitempty"` // Associated group
	// Configuración de ciclos y rangos movida a RaceDividendCycleConfig
	CycleConfigID uint                    `gorm:"column:cycle_config_id;not null" json:"cycle_config_id"`
	CycleConfig   RaceDividendCycleConfig `gorm:"foreignKey:CycleConfigID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cycle_config,omitempty"`
	Status        bool                    `gorm:"column:status;default:true" json:"status"` // Active/inactive status
	CreatedAt     time.Time               `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt     time.Time               `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt     gorm.DeletedAt          `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
	// Associated dividend ranges
	DividendRanges []RaceDividendRange `gorm:"many2many:race_dividend_config_ranges;joinForeignKey:ConfigID;joinReferences:RangeID" json:"dividend_ranges,omitempty"`
}

// TableName returns the table name for the RaceDividendConfig entity.
func (r *RaceDividendConfig) TableName() string {
	return "gaming.race_dividend_config"
}

// MarshalJSON customizes the JSON serialization of the RaceDividendConfig entity.
func (r *RaceDividendConfig) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceDividendConfig
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the RaceDividendConfig entity.
func (r *RaceDividendConfig) UnmarshalJSON(data []byte) error {
	type Alias RaceDividendConfig
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// RaceDividendRange defines payment ranges and values associated with a dividend configuration.
type RaceDividendRange struct {
	ID           uint                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string               `gorm:"column:name;type:varchar(100)" json:"name"` // Name of the range
	Configs      []RaceDividendConfig `gorm:"many2many:race_dividend_config_ranges;joinForeignKey:RangeID;joinReferences:ConfigID" json:"configs,omitempty"`
	RangoInicial float64              `gorm:"column:rango_inicial;type:decimal(14,2);not null" json:"rango_inicial"`                                 // Start range
	RangoFinal   float64              `gorm:"column:rango_final;type:decimal(14,2);not null" json:"rango_final"`                                     // End range
	TipoPago     string               `gorm:"column:tipo_pago;type:varchar(50);not null" json:"tipo_pago"`                                           // Payment type (fijo, mas, etc.)
	ValorPago    float64              `gorm:"column:valor_pago;type:decimal(14,2);not null" json:"valor_pago"`                                       // Payment value
	GroupID      *uint                `gorm:"column:group_id" json:"group_id"`                                                                       // Optional association to group
	Group        *Group               `gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group,omitempty"` // Associated group
	Status       bool                 `gorm:"column:status;default:true" json:"status"`                                                              // Active/inactive status
	CreatedAt    time.Time            `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt    time.Time            `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt    gorm.DeletedAt       `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the RaceDividendRange entity.
func (r *RaceDividendRange) TableName() string {
	return "gaming.race_dividend_range"
}

// MarshalJSON customizes the JSON serialization of the RaceDividendRange entity.
func (r *RaceDividendRange) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceDividendRange
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the RaceDividendRange entity.
func (r *RaceDividendRange) UnmarshalJSON(data []byte) error {
	type Alias RaceDividendRange
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// RaceDividend stores the official dividend per race, participant, and bet type.
type RaceDividend struct {
	ID            uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	RaceID        uint             `gorm:"column:race_id;not null" json:"race_id"`
	Race          Race             `gorm:"foreignKey:RaceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"race"`
	ParticipantID uint             `gorm:"column:participant_id;not null" json:"participant_id"`
	Participant   ParticipantsRace `gorm:"foreignKey:ParticipantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"participant"`
	TypeBetID     uint             `gorm:"column:type_bet_id;not null" json:"type_bet_id"`
	TypeBet       TypeBet          `gorm:"foreignKey:TypeBetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"type_bet"`
	Dividend      float64          `gorm:"column:dividend;type:decimal(14,6);not null" json:"dividend"`     // Official dividend value
	LoadedBy      *uint            `gorm:"column:loaded_by" json:"loaded_by"`                               // Who loaded this dividend
	LoadedAt      *time.Time       `gorm:"column:loaded_at;type:timestamp with time zone" json:"loaded_at"` // When was loaded
	Source        string           `gorm:"column:source;type:varchar(100)" json:"source"`                   // Source of the dividend data
	Status        bool             `gorm:"column:status;default:true" json:"status"`                        // Active/inactive status
	CreatedAt     time.Time        `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt     time.Time        `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt     gorm.DeletedAt   `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the RaceDividend entity.
func (r *RaceDividend) TableName() string {
	return "gaming.race_dividend"
}

// MarshalJSON customizes the JSON serialization of the RaceDividend entity.
func (r *RaceDividend) MarshalJSON() ([]byte, error) {
	if r.ID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceDividend
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the RaceDividend entity.
func (r *RaceDividend) UnmarshalJSON(data []byte) error {
	type Alias RaceDividend
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// RaceDividendConfigRange is the junction table for many-to-many relationship between RaceDividendConfig and RaceDividendRange.
type RaceDividendConfigRange struct {
	ConfigID  uint               `gorm:"primaryKey;column:config_id;not null" json:"config_id"`
	Config    RaceDividendConfig `gorm:"foreignKey:ConfigID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"config"`
	RangeID   uint               `gorm:"primaryKey;column:range_id;not null" json:"range_id"`
	Range     RaceDividendRange  `gorm:"foreignKey:RangeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"range"`
	Status    bool               `gorm:"column:status;default:true" json:"status"` // Active/inactive status for this association
	CreatedAt time.Time          `gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt time.Time          `gorm:"column:updated_at;type:timestamp with time zone;default:now()"`
	DeletedAt gorm.DeletedAt     `gorm:"index;column:deleted_at;type:timestamp with time zone" json:"-"`
}

// TableName returns the table name for the RaceDividendConfigRange entity.
func (r *RaceDividendConfigRange) TableName() string {
	return "public.race_dividend_config_ranges"
}

// MarshalJSON customizes the JSON serialization of the RaceDividendConfigRange entity.
func (r *RaceDividendConfigRange) MarshalJSON() ([]byte, error) {
	if r.ConfigID == 0 || r.RangeID == 0 {
		return []byte("null"), nil
	}

	type Alias RaceDividendConfigRange
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(r),
		},
	)
}

// UnmarshalJSON customizes the JSON deserialization of the RaceDividendConfigRange entity.
func (r *RaceDividendConfigRange) UnmarshalJSON(data []byte) error {
	type Alias RaceDividendConfigRange
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
