package initialize

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("APP_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://localhost:8080")
	})
}