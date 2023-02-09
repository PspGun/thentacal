package comment

import (
	"fmt"

	"github.com/fexcel/fexcel-backend/db"
	"github.com/fexcel/fexcel-backend/type/database"
	"github.com/fexcel/fexcel-backend/type/req"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Comment(ctx *fiber.Ctx) error {

	var body req.Comment
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	var report database.Report
	uid, err := uuid.Parse(body.ReportId)
	re := db.DB.Where("ID = ? ", uid).First(&report)
	fmt.Println(uid)
	if re.Error != nil {
		return ctx.JSON("ERORR")
	}
	db.DB.Model(report).Update("comment", body.Comment)
	return ctx.JSON(report)

}
