package main

import (
	"gametools/server/apis/project"
	"gametools/server/apis/user"
	"gametools/server/app"
	"gametools/server/models"
)

func main() {
	app.NewApp().
		InitDatabase(
			&models.User{},
			&models.Role{},
			&models.Resource{},
			&models.UserRole{},
			&models.Project{},
			&models.OperationLog{},
		).
		Api("/api",
			user.Controller{},
			project.Controller{},
		).
		Run()
}
