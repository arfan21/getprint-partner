package controllers

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/repository"
	"github.com/arfan21/getprint-partner/services"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type partnerController struct {
	service models.PartnerService
}

func NewPartnerController(db *gorm.DB, route *echo.Echo) {
	partnerRepo := repository.NewPartnerRepo(db)
	partnerService := services.NewPartnerService(partnerRepo)

	ctrl := partnerController{partnerService}

	route.POST("/partner", ctrl.Create)

}

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
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success create partner", partner))
}

func (ctrl *partnerController) Gets(c echo.Context) error {
	partners, err := ctrl.service.Gets()

	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get all partner", partners))
}

func (ctrl *partnerController) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	partner, err := ctrl.service.GetByID(uint(id))

	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get parnter", partner))
}

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
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success update partner", partner))
}
