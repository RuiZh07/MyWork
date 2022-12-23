package controller

import (
	"github.com/gofiber/fiber/v2"
)

func LoadSignUp(c *fiber.Ctx) error{
	return c.Render("signup", nil)
}