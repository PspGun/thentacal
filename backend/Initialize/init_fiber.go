package Initialize

import (
	"github.com/fexcel/fexcel-backend/enpoint"
	"github.com/gofiber/fiber/v2"
)

func Init_fiber() {
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})
	api := app.Group("/api")
	Initialize(app)
	enpoint.RegisterEnpoint(api)
	app.Listen(":8000")
}
