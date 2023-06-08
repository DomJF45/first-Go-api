package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	App         *gin.Engine
	Controllers []Controller
	Port        string
}

func (a *App) initControllers(controllers []ControllerInterface) {
	api := a.App.Group("/api")
	for _, controller := range controllers {
		controller.initRouter(api)
	}
}

func newApp(controllers []ControllerInterface, port string) App {
	var app App
	app.App = gin.Default()
	app.Port = port
	app.initControllers(controllers)
	return app
}

func (a *App) run() {
	a.App.Run(a.Port)
}
