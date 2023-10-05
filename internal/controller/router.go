package controller

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

func New() *App {
	app := App{
		echo: echo.New(),
	}

	app.initRoutes()
	return &app
}

func (a *App) initRoutes() {
	a.echo.GET("/", indexHandler)
	a.echo.GET("/healthz", checkHealthHandler)
}

func (a *App) Run() error {
	return a.echo.Start(":8080")
}
