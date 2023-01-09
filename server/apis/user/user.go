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
	app.Post("/login", login)

	app.Post("/user", addUser)
}

type ViewObject struct {
	ID        uint64          `json:"id"`
	Username  string          `json:"username"`
	Nickname  string          `json:"nickname"`
	State     models.State    `json:"state"`
	CreatedAt common.DateTime `json:"createdAt"`
	UpdatedAt common.DateTime `json:"updatedAt"`
}

// 登录
func login(c *fiber.Ctx) error {
	ctx := app.WrapContext(c)
	form := &models.User{}
	err := ctx.Request.BodyParser(form)
	if err != nil {
		return err
	}
	user := models.FindUserByUsername(form.Username)
	if user == nil {
		return ctx.Response.Error("用户名或密码错误")
	}
	formPassword := common.EncodePassword(form.Password)
	if user.Password != formPassword {
		return ctx.Response.Error("用户名或密码错误")
	}
	if user.State != models.Valid {
		return ctx.Response.Error("该用户无效")
	}
	// TODO 生成 JWT
	return ctx.Response.Ok(nil)
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
