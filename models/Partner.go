package models

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

type Partner struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   null.Time `gorm:"index" json:"deleted_at,omitempty"`
	Owner_id    uint      `gorm:"not null;unique" json:"Owner_id"`
	PartnerName string    `gorm:"size:100;not null;unique" json:"partner_name"`
	OwnerName   string    `gorm:"size:100;not null" json:"owner_name"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	PhoneNumber int64     `gorm:"size:50;not null;unique" json:"phone_number"`
	Price       Price     `gorm:"constraint:OnDelete:CASCADE;" json:"price"`
	Address     Address   `gorm:"constraint:OnDelete:CASCADE;" json:"address"`
	Status      string    `gorm:"type:enum('inactive','active');default:'inactive'" json:"status"`
}

type Price struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Print     zero.Int  `gorm:"size:30" json:"print"`
	Scan      zero.Int  `gorm:"size:30" json:"scan"`
	Fotocopy  zero.Int  `gorm:"size:30" json:"fotocopy"`
}

type Address struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  null.Time `gorm:"index" json:"deleted_at,omitempty"`
	PartnerID uint      `json:"partner_id"`
	Address   string    `gorm:"not null;size:255" json:"address"`
	Lat       string    `gorm:"not null;size:50" json:"lat"`
	Lng       string    `gorm:"not null;size:50" json:"lng"`
}

func (p Partner) Validate() error {
	return validator.ValidateStruct(&p,
		validator.Field(&p.Owner_id, validator.Required, is.Int),
		validator.Field(&p.PartnerName, validator.Required),
		validator.Field(&p.OwnerName, validator.Required),
		validator.Field(&p.Email, validator.Required, is.Email),
		validator.Field(&p.PhoneNumber, validator.Required, is.Int),
		validator.Field(&p.Address.Address, validator.Required),
		validator.Field(&p.Address.Lat, validator.Required),
		validator.Field(&p.Address.Lng, validator.Required),
	)

	return nil
}

var (
	ErrorEmailRegistered = "email already taken"
	ErrorEmailNotFound   = "email not registered"
)

type PartnerRepository interface {
	Create(partner *Partner) error
	Gets() (*[]Partner, error)
	GetByID(id uint) (*Partner, error)
	Update(id uint, partner *Partner) error
}

type PartnerService interface {
	Create(partner *Partner) error
	Gets() (*[]Partner, error)
	GetByID(id uint) (*Partner, error)
	Update(id uint, partner *Partner) error
}
