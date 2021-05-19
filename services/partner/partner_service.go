package partner

import (
	"strings"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/repository/imgur"
	_followerRepo "github.com/arfan21/getprint-partner/repository/mysql/follower"
	_partnerRepo "github.com/arfan21/getprint-partner/repository/mysql/partner"
	_userSrv "github.com/arfan21/getprint-partner/repository/user"
)

type PartnerService interface {
	Create(partner *models.Partner) error
	Fetch(name, status string) ([]*models.PartnerResponse, error)
	GetByID(id uint) (*models.PartnerResponse, error)
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
func (srv *partnerService) Fetch(name, status string) ([]*models.PartnerResponse, error) {
	var partners *[]models.Partner

	if name == "" && status != "" {
		data, err := srv.partnerRepo.Fetch("status=?", status)

		if err != nil {
			return nil, err
		}

		partners = data

	} else if status == "" && name != "" {
		data, err := srv.partnerRepo.Fetch("name LIKE ? AND status='active'", strings.ToLower("%"+name+"%"))

		if err != nil {
			return nil, err
		}

		partners = data
	} else if status == "inactive" && name != "" {
		data, err := srv.partnerRepo.Fetch("name LIKE ? AND status='inactive'", strings.ToLower("%"+name+"%"))

		if err != nil {
			return nil, err
		}

		partners = data
	} else {
		data, err := srv.partnerRepo.Fetch("status = ?", "active")

		if err != nil {
			return nil, err
		}

		partners = data
	}

	resArr := make([]*models.PartnerResponse, 0)

	for _, partner := range *partners {
		res := &models.PartnerResponse{
			ID:          partner.ID,
			CreatedAt:   partner.CreatedAt,
			UpdatedAt:   partner.UpdatedAt,
			UserID:      partner.UserID,
			Name:        partner.Name,
			Email:       partner.Email,
			PhoneNumber: partner.PhoneNumber,
			Picture:     partner.Picture,
			Address:     partner.Address.Address,
			Lat:         partner.Address.Lat,
			Lng:         partner.Address.Lng,
			Print:       partner.Price.Print,
			Scan:        partner.Price.Scan,
			Fotocopy:    partner.Price.Fotocopy,
		}
		resArr = append(resArr, res)
	}

	return resArr, nil
}

func (srv *partnerService) GetByID(id uint) (*models.PartnerResponse, error) {
	partner := new(models.Partner)
	err := srv.partnerRepo.GetByID(id, partner)
	if err != nil {
		return nil, err
	}

	totalFollower, err := srv.partnerRepo.CountFollower(id)
	if err != nil {
		return nil, err
	}

	partnerResponse := &models.PartnerResponse{
		ID:            partner.ID,
		CreatedAt:     partner.CreatedAt,
		UpdatedAt:     partner.UpdatedAt,
		UserID:        partner.UserID,
		Name:          partner.Name,
		Email:         partner.Email,
		PhoneNumber:   partner.PhoneNumber,
		Picture:       partner.Picture,
		Address:       partner.Address.Address,
		Lat:           partner.Address.Lat,
		Lng:           partner.Address.Lng,
		Print:         partner.Price.Print,
		Scan:          partner.Price.Scan,
		Fotocopy:      partner.Price.Fotocopy,
		TotalFollower: totalFollower,
	}

	return partnerResponse, nil
}

func (srv *partnerService) Update(id uint, partner *models.Partner) error {
	return nil
}
