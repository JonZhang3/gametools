package app

import (
	"fmt"
	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type RouteHandleFunc func(ctx *Context) error

type Router struct {
	app *fiber.App
}

func (r *Router) Post(path string, handler RouteHandleFunc) {
	r.app.Post(path, func(ctx *fiber.Ctx) error {
		return handler(WrapContext(ctx))
	})
}

func (r *Router) Get(path string, handler RouteHandleFunc) {
	r.app.Get(path, func(ctx *fiber.Ctx) error {
		return handler(WrapContext(ctx))
	})
}

func (r *Router) Put(path string, handler RouteHandleFunc) {
	r.app.Put(path, func(ctx *fiber.Ctx) error {
		return handler(WrapContext(ctx))
	})
}

func (r *Router) Delete(path string, handler RouteHandleFunc) {
	r.app.Delete(path, func(ctx *fiber.Ctx) error {
		return handler(WrapContext(ctx))
	})
}

type Controller interface {
	Register(router *Router)
}

const (
	resultCodeOk    = 200
	resultCodeError = 500
)

// JsonResult 接口响应
type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(data interface{}, message ...string) *JsonResult {
	msg := "操作成功"
	if len(message) > 0 {
		msg = message[0]
	}
	return &JsonResult{Code: resultCodeOk, Data: data, Message: msg}
}

func Error(message string, data ...interface{}) *JsonResult {
	var d interface{} = nil
	if len(data) > 0 {
		d = data[0]
	}
	if message == "" {
		message = "操作失败"
	}
	return &JsonResult{Code: resultCodeError, Message: message, Data: d}
}

func useJwt(app *fiber.App, cfg *config) {
	app.Use(jwtware.New(jwtware.Config{
		Filter: func(ctx *fiber.Ctx) bool {
			path := ctx.Path()
			excludePaths := cfg.Auth.Excludes
			if excludePaths != nil && len(excludePaths) > 0 {
				for _, p := range excludePaths {
					if path == p {
						return true
					}
				}
			}
			return false
		},
		SigningKey: []byte(cfg.Auth.Token.SigningKey),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			Logger.Error(err, "no JWT or JWT is expired")
			return ctx.Status(fiber.StatusUnauthorized).JSON(Error("no auth"))
		},
	}))
}

type Application struct {
	Config *config
	Router *fiber.App
}

func (app *Application) Run(controllers ...Controller) {
	server := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			message := err.Error()
			if _, ok := err.(*BizError); ok {
				return ctx.JSON(Error(message))
			}
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			Logger.Error(err, "Internal server error")
			return ctx.Status(code).JSON(Error(message))
		},
	})
	// 配置跨域
	server.Use(cors.New())
	useJwt(server, app.Config)
	// 注册 Controller
	if len(controllers) > 0 {
		router := &Router{app: server}
		for _, c := range controllers {
			if c != nil {
				c.Register(router)
			}
		}
	}
	Logger.Info("starting server")
	err := server.Listen(fmt.Sprintf(":%d", app.Config.Server.Port))
	if err != nil {
		Logger.Panic(err, "start server error")
	}
}

func (app *Application) InitDatabase(models ...any) *Application {
	dsConfig := app.Config.Datasource
	db, err := gorm.Open(mysql.Open(dsConfig.Dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		Logger.Panic(err, "init database error")
	}
	if dsConfig.AutoMigrate && len(models) > 0 {
		err = db.AutoMigrate(models...)
		if err != nil {
			Logger.Panic(err, "database auto migrate error")
		}
	}
	DB = db
	return app
}
