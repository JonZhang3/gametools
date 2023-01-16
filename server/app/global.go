package app

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	Logger = newLogger() // 日志
	DB     *gorm.DB      // 数据库 gorm
	app    *Application  // web application
)

func GetApp() *Application {
	if app == nil {
		panic(fmt.Errorf("application is not init, use NewApp() to init application"))
	}
	return app
}
