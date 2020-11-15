package main

import (
	"log"
	"os"

	"github.com/arfan21/getprint-partner/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	port := ":" + os.Getenv("PORT")
	_, err := utils.Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal(err.Error())
	}

	route := echo.New()

	route.Logger.Fatal(route.Start(port))
}
