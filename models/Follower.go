package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Follower struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `gorm:"not null;unique" json:"partner_id"`
	UserID    uint      `gorm:"not null;unique" json:"user_id"`
}

type FollowerRepository interface {
	Create(follower *Follower) error
	GetByID(id uint) (*Follower, error)
	CountFollower(partnerId uint) (int64, error)
	Delete(id uint) error
}

type FollowerService interface {
	Create(follower *Follower) error
	Delete(id uint) error
}
