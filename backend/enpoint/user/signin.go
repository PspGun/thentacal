package user

import (
	"github.com/PspGun/thentacal/db"
	"github.com/PspGun/thentacal/type/database"
	"github.com/PspGun/thentacal/type/req"
	"github.com/gofiber/fiber/v2"
)

func Signin(ctx *fiber.Ctx) error {
	var body req.Sigin
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	user := new(database.User)
	result := db.DB.Where("username = ? AND password = ?", body.Username, body.Password).First(user)
	if result.Error != nil {
		return ctx.JSON("ERORR")
	}
	return ctx.JSON(user)
}
