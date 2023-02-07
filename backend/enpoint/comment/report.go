package comment

import (
	"github.com/PspGun/thentacal/db"
	"github.com/PspGun/thentacal/type/database"
	"github.com/PspGun/thentacal/type/req"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Report(ctx *fiber.Ctx) error {

	var body req.Prompt
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	// ส่งมันไป
	// await for result
	uid, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}
	send := database.Report{
		UserId:  uid,
		Prompt:  body.Prompt,
		Result:  "select-statment",
		Comment: nil,
	}
	result := db.DB.Create(&send)
	if result.Error != nil {
		return ctx.JSON(result.Error.Error())
	}

	return ctx.JSON(send)
}
