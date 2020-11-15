package models

import (
	"time"

	"gopkg.in/guregu/null.v4/zero"
)

type Partner struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeleteAt    time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Owner_id    uint      `gorm:"not null;unique" json:"Owner_id"`
	PartnerName string    `gorm:"size:100;not null;unique" json:"partner_name"`
	OwnerName   string    `gorm:"size:100;not null" json:"owner_name"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	PhoneNumber int64     `gorm:"size:50;not null;unique" json:"phone_number"`
	Price       Price     `gorm:"constraint:OnDelete:CASCADE;" json:"price"`
	Address     Address   `gorm:"constraint:OnDelete:CASCADE;" json:"address"`
	Status      string    `gorm:"type:enum('inactive','active');default:'active'" json:"status"`
}

type Price struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Print     zero.Int  `gorm:"size:30" json:"print"`
	Scan      zero.Int  `gorm:"size:30" json:"scan"`
	Fotocopy  zero.Int  `gorm:"size:30" json:"fotocopy"`
}

type Address struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Address   string    `gorm:"size:255" json:"address"`
	Lat       string    `gorm:"size:50" json:"lat"`
	Lng       string    `gorm:"size:50" json:"lng"`
}

var (
	ErrorEmailRegistered = "email already taken"
	ErrorEmailNotFound   = "email not registered"
)

type PartnerRepository interface {
	Create(partner *Partner) error
	Get(partners *[]Partner) error
	GetByID(id uint, partner *Partner) error
	Update(partner *Partner) error
}

type PartnerService interface {
	Create(partner *Partner) error
	Get(partners *[]Partner) error
	GetByID(id uint, partner *Partner) error
	Update(partner *Partner) error
}
