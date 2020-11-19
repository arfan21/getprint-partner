package services

import (
	"github.com/arfan21/getprint-partner/models"
)

type partnerService struct {
	repo models.PartnerRepository
}

func NewPartnerService(repo models.PartnerRepository) models.PartnerService {
	services := partnerService{repo: repo}

	return &services
}

func (s *partnerService) Create(partner *models.Partner) error {
	err := s.repo.Create(partner)

	if err != nil {
		return err
	}

	return nil
}
func (s *partnerService) Gets() (*[]models.Partner, error) {
	partners, err := s.repo.Gets()

	if err != nil {
		return nil, err
	}

	return partners, nil
}
func (s *partnerService) GetByID(id uint) (*models.Partner, error) {
	partner, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return partner, nil
}
func (s *partnerService) Update(id uint, partner *models.Partner) error {
	err := s.repo.Update(id, partner)

	if err != nil {
		return err
	}

	return nil
}
