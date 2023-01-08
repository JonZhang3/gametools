package login

import (
    "gametools/server/app"
    "github.com/gofiber/fiber/v2"
)

type Controller struct {
}

func (Controller) Register(app *fiber.App) {
    app.Post("/login", login)
}

type LoginForm struct {
    Username string `form:"username"`
    Password string `form:"password"`
}

func login(c *fiber.Ctx) error {
    ctx := app.WrapContext(c)
    var lf = &LoginForm{}
    err := ctx.Request.BodyParser(lf)
    if err != nil {
        return err
    }

    return nil
}
