package repository

import (
	"github.com/arfan21/getprint-partner/models"
	"gorm.io/gorm"
)

type followerRepo struct {
	db *gorm.DB
}

func NewFollowerRepo(db *gorm.DB) models.FollowerRepository {
	return &followerRepo{db}
}

func (repo *followerRepo) Create(follower *models.Follower) error {
	err := repo.db.Create(follower).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *followerRepo) CountFollower(partnerId uint) (int64, error) {
	var count int64
	err := repo.db.Model(&models.Follower{}).Where("partner_id = ?", partnerId).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *followerRepo) GetByID(id uint) (*models.Follower, error) {
	follower := new(models.Follower)

	err := repo.db.First(&follower, id).Error

	if err != nil {
		return nil, err
	}

	return follower, nil
}

func (repo *followerRepo) Delete(id uint) error {
	err := repo.db.Unscoped().Delete(&models.Follower{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
