package follower

import (
	"sync"

	"github.com/arfan21/getprint-partner/models"
	_followerRepo "github.com/arfan21/getprint-partner/repository/mysql/follower"
	_partnerRepo "github.com/arfan21/getprint-partner/repository/mysql/partner"
	_userSrv "github.com/arfan21/getprint-partner/repository/user"
)

type FollowerService interface {
	Create(follower *models.Follower) error
	Delete(id uint) error
}

type followerService struct {
	partnerRepo  _partnerRepo.PartnerRepository
	followerRepo _followerRepo.FollowerRepository
	userSrv      _userSrv.UserServices
}

func NewFollowerService(partnerRepo _partnerRepo.PartnerRepository, followerRepo _followerRepo.FollowerRepository) FollowerService {
	userSrv := _userSrv.NewUserServices()
	return &followerService{partnerRepo, followerRepo, userSrv}
}

func (srv *followerService) Create(follower *models.Follower) error {
	var wg sync.WaitGroup

	wg.Add(2)

	chanUserError := make(chan error)
	chanPartner := make(chan *models.Partner)

	go func() {
		defer wg.Done()

		err := srv.partnerRepo.GetByID(follower.PartnerID, <-chanPartner)

		if err != nil {
			chanPartner <- nil
			return
		}
	}()

	go func() {
		defer wg.Done()

		err := srv.userSrv.VerificationUser(follower.UserID.String())

		if err != nil {
			chanUserError <- err
			return
		}

		chanUserError <- nil
	}()

	errorUser := <-chanUserError
	dataPartner := <-chanPartner

	wg.Wait()

	if dataPartner == nil || errorUser != nil {
		return models.ErrNotFound
	}

	err := srv.followerRepo.Create(follower)

	if err != nil {
		return err
	}

	return nil
}

func (service *followerService) Delete(id uint) error {
	_, err := service.followerRepo.GetByID(id)

	if err != nil {
		return err
	}

	err = service.followerRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
