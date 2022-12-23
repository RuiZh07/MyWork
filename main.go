package main

import (
	"log"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
	"NFC_Tag_UPoint/controller"
)

func main(){

	// Initialize standard go html template engine
	engine := html.New("./templates", ".html")
	
	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/", index)
	app.Get("/signup", controller.LoadSignUp)
	app.Get("/login", controller.LoadLoginPage)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
