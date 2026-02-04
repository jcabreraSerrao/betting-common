package sql

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Config struct {
	Key     string `gorm:"column:key" json:"key"`
	Value   string `gorm:"column:value" json:"value"`
	GroupId uint   `gorm:"column:group_id" json:"groupId"`
	Group   Group  `gorm:"foreignKey:GroupId;references:ID"`
	gorm.Model
}

func (c *Config) TableName() string {
	return "config.config"
}

func (c *Config) MarshalJSON() ([]byte, error) {
	if c.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Config
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(c),
		},
	)
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type Alias Config
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
