package follower

import (
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
	err := srv.userSrv.VerificationUser(follower.UserID.String())
	if err != nil {
		return models.ErrNotFound
	}

	dataPartner := new(models.Partner)
	err = srv.partnerRepo.GetByID(follower.PartnerID, dataPartner)
	if err != nil {
		return models.ErrNotFound
	}

	err = srv.followerRepo.Create(follower)

	if err != nil {
		return err
	}

	return nil
}

func (srv *followerService) Delete(id uint) error {
	_, err := srv.followerRepo.GetByID(id)

	if err != nil {
		return err
	}

	err = srv.followerRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
