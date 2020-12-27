package controllers

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/repository"
	"github.com/arfan21/getprint-partner/services"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type followerController struct {
	service models.FollowerService
}

//NewFollowerController ....
func NewFollowerController(db *gorm.DB, route *echo.Echo) {
	followerRepo := repository.NewFollowerRepo(db)
	partnerRepo := repository.NewPartnerRepo(db)
	followerService := services.NewFollowerService(followerRepo, partnerRepo)

	ctrl := followerController{followerService}

	route.POST("/follow", ctrl.Create)
	route.DELETE("/follow/:id", ctrl.Delete)
}

//Create ....
func (ctrl *followerController) Create(c echo.Context) error {
	follower := new(models.Follower)

	if err := c.Bind(follower); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err := ctrl.service.Create(follower)

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

	err = ctrl.service.Delete(uint(id))

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success unfollow partner", nil))
}
