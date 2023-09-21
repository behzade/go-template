package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("template project")
}

func checkHealthHandler(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
