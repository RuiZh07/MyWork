package controller

import (
	// "database/sql"
	// _ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	
)

func LoadDashboard(c *fiber.Ctx) error {
	return c.Render("dashboard", nil)
}