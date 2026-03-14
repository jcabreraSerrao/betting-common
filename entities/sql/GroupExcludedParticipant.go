package sql

import (
	"time"

	"gorm.io/gorm"
)

type ParticipantRoleType string

const (
	ParticipantRoleObserver ParticipantRoleType = "OBSERVER"
	ParticipantRoleMediator ParticipantRoleType = "MEDIATOR"
	ParticipantRoleOwner    ParticipantRoleType = "OWNER"
	ParticipantRolePlanner  ParticipantRoleType = "PLANNER"
	ParticipantRoleOther    ParticipantRoleType = "OTHER"
)

type GroupExcludedParticipant struct {
	ID          uint64              `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID     uint                `gorm:"column:id_group;index;uniqueIndex:idx_group_jid" json:"id_group"`
	JID         string              `gorm:"column:jid;index;uniqueIndex:idx_group_jid" json:"jid"`
	PhoneNumber string              `gorm:"column:phone_number" json:"phone_number"`
	Name        string              `gorm:"column:name" json:"name"`
	Role        ParticipantRoleType `gorm:"column:role;default:OBSERVER" json:"role"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `gorm:"index" json:"-"`
}

func (GroupExcludedParticipant) TableName() string {
	return "security.group_excluded_participants"
}
