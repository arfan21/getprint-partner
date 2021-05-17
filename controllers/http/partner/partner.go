package partner

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-partner/models"
	_followerRepo "github.com/arfan21/getprint-partner/repository/mysql/follower"
	_partnerRepo "github.com/arfan21/getprint-partner/repository/mysql/partner"
	_partnerSrv "github.com/arfan21/getprint-partner/services/partner"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/arfan21/getprint-partner/validation"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PartnerController interface {
	Routes(route *echo.Echo)
}

type partnerController struct {
	partnerSrv _partnerSrv.PartnerService
}

//NewPartnerController ...
func NewPartnerController(db *gorm.DB) PartnerController {
	partnerRepo := _partnerRepo.NewPartnerRepo(db)
	followerRepo := _followerRepo.NewFollowerRepo(db)
	partnerService := _partnerSrv.NewPartnerService(partnerRepo, followerRepo)

	return &partnerController{partnerService}
}

//Create ....
func (ctrl *partnerController) Create(c echo.Context) error {
	partner := new(models.Partner)

	if err := c.Bind(partner); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	if err := validation.Validate(*partner); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err, nil))
	}

	err := ctrl.partnerSrv.Create(partner)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success create partner", partner))
}

//Fetch .....
func (ctrl *partnerController) Fetch(c echo.Context) error {
	partners, err := ctrl.partnerSrv.Fetch(c.QueryParam("name"), c.QueryParam("status"))

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get all partner", partners))
}

//GetByID ....
func (ctrl *partnerController) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	partner, err := ctrl.partnerSrv.GetByID(uint(id))

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get partner", partner))
}

//Update ......
func (ctrl *partnerController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	partner := new(models.Partner)

	if err := c.Bind(partner); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	if err := validation.Validate(*partner); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err, nil))
	}

	err = ctrl.partnerSrv.Update(uint(id), partner)

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success update partner", partner))
}
