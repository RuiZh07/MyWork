package controller

import (
	// "log"
	"github.com/gofiber/fiber/v2"
)

// Render setting.html template
func LoadSettingPage(c *fiber.Ctx) error {
	return c.Render("setting", nil)
}

func LoadChangeUsername(c *fiber.Ctx) error{
	return c.Render("changeUsername", nil)
}

func LoadChangePassword(c *fiber.Ctx) error {
	return c.Render("changePassword", nil)
}

