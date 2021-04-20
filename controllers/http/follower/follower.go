package follower

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-partner/models"
	_followerRepo "github.com/arfan21/getprint-partner/repository/mysql/follower"
	_partnerRepo "github.com/arfan21/getprint-partner/repository/mysql/partner"
	_followerSrv "github.com/arfan21/getprint-partner/services/follower"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PartnerController interface {
	Routes(route *echo.Echo)
}

type followerController struct {
	followerSrv _followerSrv.FollowerService
}

//NewFollowerController ....
func NewFollowerController(db *gorm.DB) PartnerController {
	followerRepo := _followerRepo.NewFollowerRepo(db)
	partnerRepo := _partnerRepo.NewPartnerRepo(db)
	followerService := _followerSrv.NewFollowerService(partnerRepo, followerRepo)

	return &followerController{followerService}
}

//Create ....
func (ctrl *followerController) Create(c echo.Context) error {
	follower := new(models.Follower)

	if err := c.Bind(follower); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err := ctrl.followerSrv.Create(follower)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success follow partner", follower))
}

//Delete ....
func (ctrl *followerController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err = ctrl.followerSrv.Delete(uint(id))

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success unfollow partner", nil))
}
