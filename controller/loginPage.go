package controller

import (
	"github.com/gofiber/fiber/v2"
)

func LoadLoginPage(c *fiber.Ctx) error {
	// Render login.html template
	return c.Render("login", nil)
}
