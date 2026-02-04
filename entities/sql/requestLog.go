package sql

import "time"

type RequestLog struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Method    string    `gorm:"size:10;not null"`
	URL       string    `gorm:"type:text;not null"`
	Body      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (RequestLog) TableName() string {
	return "security.request_log"
}
