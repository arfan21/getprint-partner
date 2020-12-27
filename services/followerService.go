package services

import (
	"sync"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/utils"
)

type followerService struct {
	repo        models.FollowerRepository
	repoPartner models.PartnerRepository
}

func NewFollowerService(repo models.FollowerRepository, repoPartner models.PartnerRepository) models.FollowerService {
	return &followerService{repo, repoPartner}
}

func (service *followerService) Create(follower *models.Follower) error {
	var wg sync.WaitGroup

	wg.Add(2)

	chanUser := make(chan map[string]interface{})
	chanPartner := make(chan *models.Partner)

	go func() {
		defer wg.Done()

		data, err := service.repoPartner.GetByID(follower.PartnerID)

		if err != nil {
			chanPartner <- nil
			return
		}

		chanPartner <- data
	}()

	go func() {
		defer wg.Done()

		data, err := utils.GetUser(follower.UserID)

		if err != nil {
			chanUser <- nil
			return
		}

		chanUser <- data
	}()

	dataUser := <-chanUser
	dataPartner := <-chanPartner

	wg.Wait()

	if dataPartner == nil || len(dataUser) == 0 {
		return models.ErrNotFound
	}

	err := service.repo.Create(follower)

	if err != nil {
		return err
	}

	return nil
}

func (service *followerService) Delete(id uint) error {
	_, err := service.repo.GetByID(id)

	if err != nil {
		return err
	}

	err = service.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
