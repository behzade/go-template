package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	return c.JSON(
		map[string]string{
			"service_name": "template",
		})
}

func checkHealthHandler(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
