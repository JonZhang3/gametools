package controllers

import (
	"gametools/server/app"
	"gametools/server/common"
	"gametools/server/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
}

func (UserController) Register(app *fiber.App) {
	app.Post("/usr", addUser)
}

type UserVO struct {
	ID        uint64          `json:"id"`
	Username  string          `json:"username"`
	Nickname  string          `json:"nickname"`
	State     models.State    `json:"state"`
	CreatedAt common.DateTime `json:"createdAt"`
	UpdatedAt common.DateTime `json:"updatedAt"`
}

func addUser(c *fiber.Ctx) error {
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		return err
	}
	app.DB.Create(user)
	return c.JSON(app.Ok(user))
}
