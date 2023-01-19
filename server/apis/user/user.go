package user

import (
	"gametools/server/app"
	"gametools/server/common"
	"gametools/server/models"
)

type Controller struct {
}

func (Controller) Register(router *app.Router) {
	router.Post("/login", login)

	router.Post("/user", addUser)
}

type ViewObject struct {
	ID        uint64          `json:"id"`
	Username  string          `json:"username"`
	Nickname  string          `json:"nickname"`
	State     models.State    `json:"state,omitempty"`
	CreatedAt common.DateTime `json:"createdAt,omitempty"`
	UpdatedAt common.DateTime `json:"updatedAt,omitempty"`
}

// 登录
func login(ctx *app.Context) error {
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
	token, err := app.NewToken(&app.TokenClaims{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		IsAdmin:  user.IsAdmin,
	})
	if err != nil {
		return err
	}
	return ctx.Response.Ok(map[string]any{
		"token": token,
		"user": &ViewObject{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: common.DateTime(user.CreatedAt),
			UpdatedAt: common.DateTime(user.UpdatedAt),
		},
	})
}

// 新增用户
func addUser(ctx *app.Context) error {
	user := &models.User{}
	if err := ctx.Request.BodyParser(user); err != nil {
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
	return ctx.Response.Ok(vo)
}
