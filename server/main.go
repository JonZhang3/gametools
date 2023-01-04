package main

import (
	"gametools/server/controllers"
	app "gametools/server/framework"
)

func main() {
	app.NewApp().
		Run(
			controllers.UserController{},
		)
}
