package main

import (
	"fmt"
	"log"
	"os"

	_followerCtrl "github.com/arfan21/getprint-partner/controllers/http/follower"
	_partnerCtrl "github.com/arfan21/getprint-partner/controllers/http/partner"
	"github.com/arfan21/getprint-partner/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())

	followerCtrl := _followerCtrl.NewFollowerController(db)
	followerCtrl.Routes(route)

	partnerCtrl := _partnerCtrl.NewPartnerController(db)
	partnerCtrl.Routes(route)

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
