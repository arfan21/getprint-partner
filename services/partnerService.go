package services

import (
	"strings"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/labstack/echo/v4"
)

type partnerService struct {
	repo         models.PartnerRepository
	repoFollower models.FollowerRepository
}

//NewPartnerService ...
func NewPartnerService(repo models.PartnerRepository, repoFollower models.FollowerRepository) models.PartnerService {
	services := partnerService{repo, repoFollower}

	return &services
}

//Create ....
func (s *partnerService) Create(partner *models.Partner) error {
	_, err := utils.GetUser(partner.UserID)

	if err != nil {
		return err
	}

	err = s.repo.Create(partner)

	if err != nil {
		return err
	}

	return nil
}

//Fetch ....
func (s *partnerService) Fetch(c echo.Context) (*[]models.Partner, error) {
	q := c.QueryParam("q")
	status := c.QueryParam("status")

	if q == "" {
		partners, err := s.repo.Fetch("status=?", status)

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else if status == "" {
		partners, err := s.repo.Fetch("name LIKE ? AND status='active'", strings.ToLower("%"+q+"%"))

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else if status == "inactive" && q != "" {
		partners, err := s.repo.Fetch("name LIKE ? AND status='inactive'", strings.ToLower("%"+q+"%"))

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else {
		partners, err := s.repo.Fetch("status = ?", "active")

		if err != nil {
			return nil, err
		}

		return partners, nil
	}

}

//GetByID ...
func (s *partnerService) GetByID(id uint) (*models.PartnerWithCountFollower, error) {
	partner, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}
	count, err := s.repoFollower.CountFollower(id)

	if err != nil {
		return nil, err
	}

	return &models.PartnerWithCountFollower{
		Partner:       *partner,
		TotalFollower: count,
	}, nil
}

//Update ....
func (s *partnerService) Update(id uint, partner *models.Partner) error {
	if partner.Status == "" {
		partner.Status = "inactive"
	}

	_, err := s.repo.GetByID(id)

	if err != nil {
		return err
	}

	err = s.repo.Update(id, partner)

	if err != nil {
		return err
	}

	return nil
}
