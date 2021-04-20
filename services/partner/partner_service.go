package partner

import (
	"strings"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/repository/imgur"
	_followerRepo "github.com/arfan21/getprint-partner/repository/mysql/follower"
	_partnerRepo "github.com/arfan21/getprint-partner/repository/mysql/partner"
	_userSrv "github.com/arfan21/getprint-partner/repository/user"
	"github.com/labstack/echo/v4"
)

type PartnerService interface {
	Create(partner *models.Partner) error
	Fetch(c echo.Context) (*[]models.Partner, error)
	GetByID(id uint) (*models.PartnerWithCountFollower, error)
	Update(id uint, partner *models.Partner) error
}

type partnerService struct {
	partnerRepo  _partnerRepo.PartnerRepository
	repoFollower _followerRepo.FollowerRepository
	userSrv      _userSrv.UserServices
	imgurSrv     imgur.Imgur
}

//NewPartnerService ...
func NewPartnerService(partnerRepo _partnerRepo.PartnerRepository, followerRepo _followerRepo.FollowerRepository) PartnerService {
	userSrv := _userSrv.NewUserServices()
	imgurSrv := imgur.NewImgur()
	return &partnerService{partnerRepo, followerRepo, userSrv, imgurSrv}
}

//Create ....
func (srv *partnerService) Create(partner *models.Partner) error {
	err := srv.userSrv.VerificationUser(partner.UserID.String())

	if err != nil {
		return err
	}

	res, err := srv.imgurSrv.Upload(partner.Picture)

	if err != nil {
		return err
	}

	partner.Picture = res.Data.Link
	partner.DeleteHash = res.Data.DeleteHash

	err = srv.partnerRepo.Create(partner)

	if err != nil {
		errDelete := srv.imgurSrv.Delete(partner.DeleteHash)
		if errDelete != nil {
			err = models.ErrInternalServerError
		}
		return err
	}
	partner.DeleteHash = ""

	return nil
}

//Fetch ....
func (srv *partnerService) Fetch(c echo.Context) (*[]models.Partner, error) {
	q := c.QueryParam("q")
	status := c.QueryParam("status")

	if q == "" {
		partners, err := srv.partnerRepo.Fetch("status=?", status)

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else if status == "" {
		partners, err := srv.partnerRepo.Fetch("name LIKE ? AND status='active'", strings.ToLower("%"+q+"%"))

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else if status == "inactive" && q != "" {
		partners, err := srv.partnerRepo.Fetch("name LIKE ? AND status='inactive'", strings.ToLower("%"+q+"%"))

		if err != nil {
			return nil, err
		}

		return partners, nil
	} else {
		partners, err := srv.partnerRepo.Fetch("status = ?", "active")

		if err != nil {
			return nil, err
		}

		return partners, nil
	}

}

func (srv *partnerService) GetByID(id uint) (*models.PartnerWithCountFollower, error) {
	return nil, nil
}

func (srv *partnerService) Update(id uint, partner *models.Partner) error {
	return nil
}
