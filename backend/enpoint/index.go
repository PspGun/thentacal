package enpoint

import (
	"github.com/PspGun/thentacal/enpoint/comment"
	"github.com/PspGun/thentacal/enpoint/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterEnpoint(group fiber.Router) {
	group.Get("/signin", user.Signin)
	group.Post("/signup", user.Signup)
	group.Post("/comment", comment.Comment)
	group.Post("/prompt", comment.Report)
}
