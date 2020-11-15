package repository

import (
	"github.com/arfan21/getprint-partner/models"
	"gorm.io/gorm"
)

type mysqlPartnerRepo struct {
	db *gorm.DB
}

func NewPartnerRepo(db *gorm.DB) models.PartnerRepository {
	repo := mysqlPartnerRepo{db}

	return &repo
}

func (repo *mysqlPartnerRepo) Create(partner *models.Partner) error
func (repo *mysqlPartnerRepo) Get(partners *[]models.Partner) error
func (repo *mysqlPartnerRepo) GetByID(id uint, partner *models.Partner) error
func (repo *mysqlPartnerRepo) Update(partner *models.Partner) error
