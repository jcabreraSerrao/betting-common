package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	Name          string `gorm:"column:name"`
	LastName      string `gorm:"column:last_name"`
	Email         string `gorm:"column:email"`
	User          string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	RecoveryToken string `gorm:"column:recovery_token"`
	Status        bool   `gorm:"column:status"`
	NameSearch    string `gorm:"column:name_search"`
	PhoneNumber   string `gorm:"column:phone_number"`
	gorm.Model
}

func (s *User) TableName() string {
	return "security.users"
}

func (s *User) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias User
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
			// Note the use of `(*Alias)(&r)`
		},
	)
}
