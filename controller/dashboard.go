package controller

import (
	// "database/sql"
	// _ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	
)

// func init() {
// 	var err error

// 	db, err = sql.Open("postgres", "postgres://admin:admin@localhost:5432/wacave?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func LoadDashboard(c *fiber.Ctx) error {
	// err := db.QueryRow("SELECT name FROM universities WHERE domain = $1", domain).Scan(&university)
	// if err != nil {
	// 	return err
	// }
	return c.Render("dashboard", nil)
}