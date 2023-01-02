package framework

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

type Controller interface {
    Register(app *fiber.App)
}

const (
    resultCodeOk    = 200
    resultCodeError = 501
)

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

type application struct {
}

func (*application) Run(controllers ...Controller) {
    server := fiber.New(fiber.Config{
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            message := ""
            if e, ok := err.(*BizError); ok {
                message = e.Error()
                return ctx.JSON(Error(message))
            }
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
                message = e.Error()
                Logger.Error(e, "")
            }
            return ctx.Status(code).JSON(Error(message))
        },
    })
    // 配置跨域
    server.Use(cors.New())
    // 注册Controller
    if controllers != nil && len(controllers) > 0 {
        for _, c := range controllers {
            if c != nil {
                c.Register(server)
            }
        }
    }
    Logger.Info("starting server")
    err := server.Listen("")
    if err != nil {
        panic(fmt.Errorf("start server error %v", err))
    }
    
}
