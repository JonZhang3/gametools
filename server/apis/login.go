package apis

import "github.com/gofiber/fiber/v2"

type LoginController struct {
}

func (LoginController) Register(app *fiber.App) {
	app.Post("/login", login)
}

func login(ctx *fiber.Ctx) error {

	return nil
}
