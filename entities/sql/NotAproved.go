package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type FlagNotAproved struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	// Refills []Refills `gorm:"foreignKey:FlagNotAprovedId;references:ID"`
}

func (b *FlagNotAproved) TableName() string {
	return "transactions.FlagNotApproved"
}

func (b *FlagNotAproved) MarshalJSON() ([]byte, error) {
	if b.ID == 0 {
		return []byte("null"), nil
	}

	type Alias FlagNotAproved
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(b),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
