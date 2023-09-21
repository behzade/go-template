package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber fiber.App
}

func New() *App {
	app := App{
		fiber: *fiber.New(),
	}

	app.initRoutes()
	return &app
}

func (a *App) initRoutes() {
	a.fiber.Get("/", indexHandler)
	a.fiber.Get("/healthz", checkHealthHandler)
}

func (a *App) Run() {
	log.Fatal(a.fiber.Listen(":8080"))
}
