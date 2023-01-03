package main

import (
	"NFC_Tag_UPoint/controller"
	"NFC_Tag_UPoint/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"time"
)

func main() {

	// log.Println("Waitting for server to boot in 10s")
	time.Sleep(10 * time.Second)

	// Create table in database
	database.CreateTable()

	// Load University Data into Database
	database.LoadUniversityData()

	// Initialize standard go html template engine
	engine := html.New("./templates", ".html")

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/", index)
	app.Get("/signup", controller.LoadRegister)
	app.Get("/login", controller.LoadLoginPage)
	app.Post("/selectU", controller.HandleUniversitySelection)
	app.Post("/register", controller.HandleRegistration)
	app.Post("/login", controller.HandleLogin)
	app.Get("/dashboard", controller.LoadDashboard)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
