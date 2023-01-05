package main

import (
	"gametools/server/app"
	"gametools/server/controllers"
	"gametools/server/models"
)

func main() {
	app.NewApp().
		InitDatabase(
            &models.User{},
            &models.Role{},
        ).
		Run(
			controllers.UserController{},
		)
}
