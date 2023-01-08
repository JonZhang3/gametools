package main

import (
    "gametools/server/apis"
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
        Run(
            apis.UserController{},
        )
}
