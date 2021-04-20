package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type Follower struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `gorm:"not null;unique" json:"partner_id"`
	UserID    uuid.UUID `gorm:"not null;unique;type:char(36)" json:"user_id"`
}
