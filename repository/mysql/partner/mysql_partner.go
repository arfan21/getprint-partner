package partner

import (
	"time"

	"github.com/arfan21/getprint-partner/models"
	"gorm.io/gorm"
)

type PartnerRepository interface {
	Create(partner *models.Partner) error
	Fetch(query string, args string) (*[]models.Partner, error)
	GetByID(id uint, partner *models.Partner) error
	Update(id string, partner *models.Partner) error
	CountFollower(idPartner uint) (int64, error)
}

type mysqlPartnerRepo struct {
	db *gorm.DB
}

//NewPartnerRepo ....
func NewPartnerRepo(db *gorm.DB) PartnerRepository {
	return &mysqlPartnerRepo{db}
}

//Create ....
func (repo *mysqlPartnerRepo) Create(partner *models.Partner) error {
	return repo.db.Create(partner).Error
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
func (repo *mysqlPartnerRepo) GetByID(id uint, partner *models.Partner) error {
	return repo.db.Preload("Price").Preload("Address").First(&partner, id).Error
}

//Update ....
func (repo *mysqlPartnerRepo) Update(id string, partner *models.Partner) error {

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

	return repo.db.Debug().Table("addresses").Where("partner_id = ?", id).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"address":    partner.Address.Address,
		"lat":        partner.Address.Lat,
		"lng":        partner.Address.Lng,
	}).Error
}

func (repo *mysqlPartnerRepo) CountFollower(idPartner uint) (int64, error) {
	follower := new(models.Follower)
	var count int64
	err := repo.db.Model(follower).Where("partner_id = ?", idPartner).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
