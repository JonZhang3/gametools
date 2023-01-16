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

func NewApp() *Application {
	if app != nil {
		panic("the application already initialized")
	}
	a := &Application{
		Config: &config{},
	}
	initConfig(a)
	app = a
	return a
}
