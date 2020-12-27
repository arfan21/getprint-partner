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

type partnerController struct {
	service models.PartnerService
}

//NewPartnerController ...
func NewPartnerController(db *gorm.DB, route *echo.Echo) {
	partnerRepo := repository.NewPartnerRepo(db)
	followerRepo := repository.NewFollowerRepo(db)
	partnerService := services.NewPartnerService(partnerRepo, followerRepo)

	ctrl := partnerController{partnerService}

	route.POST("/partner", ctrl.Create)
	route.GET("/partner", ctrl.Fetch)
	route.GET("/partner/:id", ctrl.GetByID)
	route.PUT("/partner/:id", ctrl.Update)

}

//Create ....
func (ctrl *partnerController) Create(c echo.Context) error {
	partner := new(models.Partner)

	if err := c.Bind(partner); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	if err := partner.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", "error validating data", err))
	}

	err := ctrl.service.Create(partner)

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success create partner", partner))
}

//Fetch .....
func (ctrl *partnerController) Fetch(c echo.Context) error {
	partners, err := ctrl.service.Fetch(c)

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

	partner, err := ctrl.service.GetByID(uint(id))

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

	if err := partner.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", "error validating data", err))
	}

	err = ctrl.service.Update(uint(id), partner)

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success update partner", partner))
}
