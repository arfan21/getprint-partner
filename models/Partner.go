package models

import (
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

type Partner struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   null.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID      uint      `gorm:"not null;unique" json:"user_id"`
	Name        string    `gorm:"size:100;not null;unique" json:"name"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	PhoneNumber string    `gorm:"size:50;not null;unique" json:"phone_number"`
	Picture     string    `gorm:"not null" json:"picture"`
	DeleteHash  string    `gorm:"not null" json:"delete_hash,omitempty"`
	Price       Price     `gorm:"constraint:OnDelete:CASCADE;" json:"price"`
	Address     Address   `gorm:"constraint:OnDelete:CASCADE;" json:"address"`
	Status      string    `gorm:"type:enum('inactive','active');default:'inactive'" json:"status"`
}

type PartnerWithCountFollower struct {
	Partner
	TotalFollower int64 `json:"total_follower"`
}

type Address struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Address   string    `gorm:"not null;size:255" json:"address"`
	Lat       string    `gorm:"not null;size:50" json:"lat"`
	Lng       string    `gorm:"not null;size:50" json:"lng"`
}

type Price struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Print     zero.Int  `gorm:"size:30" json:"print"`
	Scan      zero.Int  `gorm:"size:30" json:"scan"`
	Fotocopy  zero.Int  `gorm:"size:30" json:"fotocopy"`
}

type PartnerRepository interface {
	Create(partner *Partner) error
	Fetch(query string, args string) (*[]Partner, error)
	GetByID(id uint) (*Partner, error)
	Update(id uint, partner *Partner) error
}

type PartnerService interface {
	Create(partner *Partner) error
	Fetch(c echo.Context) (*[]Partner, error)
	GetByID(id uint) (*PartnerWithCountFollower, error)
	Update(id uint, partner *Partner) error
}
