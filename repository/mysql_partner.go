package repository

import (
	"log"
	"time"

	"github.com/arfan21/getprint-partner/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mysqlPartnerRepo struct {
	db *gorm.DB
}

func NewPartnerRepo(db *gorm.DB) models.PartnerRepository {
	repo := mysqlPartnerRepo{db}

	return &repo
}

func (repo *mysqlPartnerRepo) Create(partner *models.Partner) error {
	err := repo.db.Create(partner).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *mysqlPartnerRepo) Gets() (*[]models.Partner, error) {
	partners := make([]models.Partner, 0)

	err := repo.db.Preload("Price").Preload("Address").Find(&partners).Error

	if err != nil {
		return nil, err
	}

	return &partners, nil
}
func (repo *mysqlPartnerRepo) GetByID(id uint) (*models.Partner, error) {
	partner := new(models.Partner)
	err := repo.db.Preload("Price").Preload("Address").Find(&partner, id).Error

	if err != nil {
		return nil, err
	}

	log.Println(partner.ID)

	if partner.ID == 0 {
		return nil, errors.New("partner not found")
	}

	return partner, nil
}
func (repo *mysqlPartnerRepo) Update(id uint, partner *models.Partner) error {
	partnerDB, err := repo.GetByID(id)

	if err != nil {
		return err
	}

	err = repo.db.Debug().Table("partners").Where("id = ?", partnerDB.ID).Updates(map[string]interface{}{
		"updated_at":   time.Now(),
		"partner_name": partner.PartnerName,
		"owner_name":   partner.OwnerName,
		"email":        partner.Email,
		"phone_number": partner.PhoneNumber,
		"status":       partner.Status,
	}).Error

	if err != nil {
		return err
	}

	err = repo.db.Debug().Table("prices").Where("partner_id = ?", partnerDB.ID).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"print":      partner.Price.Print,
		"fotocopy":   partner.Price.Fotocopy,
		"scan":       partner.Price.Scan,
	}).Error

	if err != nil {
		return err
	}

	err = repo.db.Debug().Table("addresses").Where("partner_id = ?", partnerDB.ID).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"address":    partner.Address.Address,
		"lat":        partner.Address.Lat,
		"lng":        partner.Address.Lng,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
