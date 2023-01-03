package controller

import (

	"github.com/gofiber/fiber/v2"
)

func LoadDashboard(c *fiber.Ctx) error {

	return c.Render("dashboard", nil)
}
