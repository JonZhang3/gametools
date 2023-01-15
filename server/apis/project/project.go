package project

import (
	"gametools/server/app"
	"gametools/server/models"
)

type Controller struct {
}

func (Controller) Register(router *app.Router) {
	router.Get("", listProjects)
	router.Post("", createProject)
}

func listProjects(ctx *app.Context) error {
	return nil
}

func createProject(ctx *app.Context) error {
	form := &models.Project{}
	err := ctx.Request.BodyParser(form)
	if err != nil {
		return err
	}
	
	return nil
}
