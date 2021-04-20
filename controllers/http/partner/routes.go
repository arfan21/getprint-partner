package partner

import "github.com/labstack/echo/v4"

func (ctrl partnerController) Routes(route *echo.Echo) {
	route.POST("/partner", ctrl.Create)
	route.GET("/partner", ctrl.Fetch)
	route.GET("/partner/:id", ctrl.GetByID)
	route.PUT("/partner/:id", ctrl.Update)
}
