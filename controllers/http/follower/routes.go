package follower

import "github.com/labstack/echo/v4"

func (ctrl followerController) Routes(route *echo.Echo) {
	route.POST("/follow", ctrl.Create)
	route.DELETE("/follow/:id", ctrl.Delete)
}
