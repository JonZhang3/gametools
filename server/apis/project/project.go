package project

import (
	"gametools/server/apis"
	"gametools/server/app"
	"gametools/server/models"
)

type Controller struct {
}

func (Controller) Register(router *app.Router) {
	router.Get("/project", pageListProjects)
	router.Post("/project", createProject)
}

type projectForm struct {
	models.Project
	apis.Query
}

func pageListProjects(ctx *app.Context) error {
	form := &projectForm{}
	err := ctx.Request.QueryParser(form)
	if err != nil {
		return err
	}
	pager := models.PageListProjects(&form.Project, form.GetOffset(), form.GetLimit())
	return ctx.Response.Ok(pager)
}

func createProject(ctx *app.Context) error {
	form := &projectForm{}
	err := ctx.Request.BodyParser(form)
	if err != nil {
		return err
	}
	existedProject := models.FindProjectByName(form.Name)
	if existedProject != nil {
		return ctx.Response.Error("项目已存在")
	}
	project := &models.Project{
		Name:        form.Name,
		Description: form.Description,
		CreatedBy:   ctx.CurrentUserId(),
	}
	result := app.DB.Create(project)
	if result.Error != nil {
		app.Logger.Error(result.Error, "create project error")
		return ctx.Response.Error(result.Error.Error())
	}
	return ctx.Response.Ok(nil)
}
