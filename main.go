package main

import (
	"log"
	"os"

	"github.com/arfan21/getprint-partner/controllers"
	"github.com/arfan21/getprint-partner/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	port := ":" + os.Getenv("PORT")
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	route := echo.New()
	controllers.NewPartnerController(db, route)
	controllers.NewFollowerController(db, route)

	route.Logger.Fatal(route.Start(port))
}
