package user

import (
    "gametools/server/app"
    "gametools/server/common"
    "gametools/server/models"

    "github.com/gofiber/fiber/v2"
)

type Controller struct {
}

func (Controller) Register(app *fiber.App) {
    app.Post("/usr", addUser)
}

type ViewObject struct {
    ID        uint64          `json:"id"`
    Username  string          `json:"username"`
    Nickname  string          `json:"nickname"`
    State     models.State    `json:"state"`
    CreatedAt common.DateTime `json:"createdAt"`
    UpdatedAt common.DateTime `json:"updatedAt"`
}

type Form struct {
    Username string `form:"username"`
    Password string `form:"password"`
}

// 登录
func login(c *fiber.Ctx) error {
    ctx := app.WrapContext(c)
    user := &models.User{}
    err := ctx.Request.BodyParser(user)
    if err != nil {
        return err
    }

    return nil
}

// 新增用户
func addUser(c *fiber.Ctx) error {
    user := &models.User{}
    if err := c.BodyParser(user); err != nil {
        return err
    }
    user.Password = common.EncodePassword(user.Password)
    user.State = models.Valid
    r := app.DB.Create(user)
    if r.Error != nil {
        return r.Error
    }
    vo := &ViewObject{
        ID:        user.ID,
        Username:  user.Username,
        Nickname:  user.Nickname,
        State:     user.State,
        CreatedAt: common.DateTime(user.CreatedAt),
        UpdatedAt: common.DateTime(user.UpdatedAt),
    }
    return c.JSON(app.Ok(vo))
}
