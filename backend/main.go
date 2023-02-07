package main

import (
	"fmt"
	"github.com/PspGun/thentacal/enpoint"
	"log"
	"os"

	"github.com/PspGun/thentacal/Initialize"
	"github.com/PspGun/thentacal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if ok := os.Getenv("GO_ENV"); ok != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})
	err := db.DBsetup()
	if err != nil {
		fmt.Println(err)
		return
	}
	api := app.Group("/api")
	Initialize.Initialize(app)
	enpoint.RegisterEnpoint(api)
	app.Listen(":8000")

}
