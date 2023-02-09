package user

import (
	"github.com/fexcel/fexcel-backend/db"
	"github.com/fexcel/fexcel-backend/type/database"
	"github.com/fexcel/fexcel-backend/type/req"
	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {

	var body req.Signup
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	user := database.User{
		Email:    body.Email,
		Username: body.Username,
		Password: body.Password,
	}

	db.DB.Create(&user)
	send := database.User{
		Email:    body.Email,
		Username: body.Username,
	}
	return ctx.JSON(send)
}
