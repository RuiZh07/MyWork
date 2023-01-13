package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func routeNotExist(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		c.SendStatus(404)
		c.SendString("404 Not Found")
	}
	return nil
}
