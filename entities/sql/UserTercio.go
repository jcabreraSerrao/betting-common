package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type UserTercio struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Name      string    `gorm:"column:name" json:"name"`
	Status    bool      `gorm:"column:status" json:"status"`
	Slug      string    `gorm:"column:slug" json:"slug"`
	Password  string    `gorm:"column:password" json:"-"`
	Token     string    `gorm:"column:token" json:"-"`
	Tercios   []Tercios `gorm:"foreignKey:UserTercioID;references:ID" json:"tercios"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *UserTercio) TableName() string {
	return "security.user_tercio"
}

func (u *UserTercio) MarshalJSON() ([]byte, error) {
	if u.ID == 0 {
		return []byte("null"), nil
	}

	type Alias UserTercio
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(u),
		},
	)
}

func (u *UserTercio) UnmarshalJSON(data []byte) error {
	type Alias UserTercio
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
