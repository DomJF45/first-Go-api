package main

import (
	"github.com/gin-gonic/gin"

	"github.com/web-service-gin/src/interfaces"
)

type App struct {
	App         *gin.Engine
	Controllers []interfaces.ControllerInterface
	Port        string
}

func (a *App) initControllers(controllers []interfaces.ControllerInterface) {
	api := a.App.Group("/api")
	for _, controller := range controllers {
		controller.InitRouter(api)
	}
}

func newApp(controllers []interfaces.ControllerInterface, port string) App {
	var app App
	app.App = gin.Default()
	app.Port = port
	app.initControllers(controllers)
	return app
}

func (a *App) run() {
	a.App.Run(a.Port)
}
