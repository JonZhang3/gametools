package controllers

import "github.com/gofiber/fiber/v2"

type UserController struct {
}

func (UserController) Register(app *fiber.App) {
	app.Get("/usr", list)
}

func list(c *fiber.Ctx) error {
	return nil
}
