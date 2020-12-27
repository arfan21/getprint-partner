package repository

import (
	"time"

	"github.com/arfan21/getprint-partner/models"
	"gorm.io/gorm"
)

type mysqlPartnerRepo struct {
	db *gorm.DB
}

//NewPartnerRepo ....
func NewPartnerRepo(db *gorm.DB) models.PartnerRepository {

	return &mysqlPartnerRepo{db}
}

//Create ....
func (repo *mysqlPartnerRepo) Create(partner *models.Partner) error {
	err := repo.db.Create(partner).Error

	if err != nil {
		return err
	}

	return nil
}

//Fetch ....
func (repo *mysqlPartnerRepo) Fetch(query string, args string) (*[]models.Partner, error) {
	partners := make([]models.Partner, 0)

	err := repo.db.Debug().Preload("Price").Preload("Address").Where(query, args).Find(&partners).Error

	if err != nil {
		return nil, err
	}

	return &partners, nil
}

//GetByID ....
func (repo *mysqlPartnerRepo) GetByID(id uint) (*models.Partner, error) {
	partner := new(models.Partner)
	err := repo.db.Preload("Price").Preload("Address").First(&partner, id).Error

	if err != nil {
		return nil, err
	}

	return partner, nil
}

//Update ....
func (repo *mysqlPartnerRepo) Update(id uint, partner *models.Partner) error {

	err := repo.db.Debug().Table("partners").Where("id = ?", id).Updates(map[string]interface{}{
		"updated_at":   time.Now(),
		"name":         partner.Name,
		"email":        partner.Email,
		"phone_number": partner.PhoneNumber,
		"status":       partner.Status,
	}).Error

	err = repo.db.Debug().Table("prices").Where("partner_id = ?", id).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"print":      partner.Price.Print,
		"fotocopy":   partner.Price.Fotocopy,
		"scan":       partner.Price.Scan,
	}).Error

	if err != nil {
		return err
	}

	err = repo.db.Debug().Table("addresses").Where("partner_id = ?", id).Updates(map[string]interface{}{
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
