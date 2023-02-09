package user

import (
	"github.com/fexcel/fexcel-backend/db"
	"github.com/fexcel/fexcel-backend/type/database"
	"github.com/fexcel/fexcel-backend/type/req"
	"github.com/fexcel/fexcel-backend/utill"
	"github.com/gofiber/fiber/v2"
)

func Signin(ctx *fiber.Ctx) error {
	var body req.Sigin
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	user := new(database.User)
	if utill.Valid(body.Username) {
		result := db.DB.Where("email = ? AND password = ?", body.Username, body.Password).First(user)
		if result.Error != nil {
			return ctx.JSON("ERORR")
		}
		return ctx.JSON(user)
	} else {
		result := db.DB.Where("username = ? AND password = ?", body.Username, body.Password).First(user)
		if result.Error != nil {
			return ctx.JSON("ERORR")
		}
		return ctx.JSON(user)
	}
}
