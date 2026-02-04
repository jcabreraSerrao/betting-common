package sql

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type LogsGroup struct {
	Data       string    `gorm:"type:text;column:data"`
	Code       string    `gorm:"type:varchar;column:code"`
	Template   string    `gorm:"type:text;column:template"`
	Date       time.Time `gorm:"type:timestamp without time zone;column:date"`
	IP         string    `gorm:"type:varchar;column:ip"`
	Before     string    `gorm:"type:text;column:before"`
	After      string    `gorm:"type:text;column:after"`
	MakerID    int       `gorm:"foreignKey:MakerID;associationForeignKey:ID"`
	Maker      UserGroup `gorm:"foreignKey:MakerID;references:ID"`
	ReceiverID int       `gorm:"foreignKey:ReceiverID;associationForeignKey:ID"`
	Receiver   UserGroup `gorm:"foreignKey:ReceiverID;references:ID"`
	gorm.Model
}

func (h *LogsGroup) TableName() string {
	return "config.logs_group"
}

func (h *LogsGroup) MarshalJSON() ([]byte, error) {
	type Alias LogsGroup
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(h),
		},
	)
}
