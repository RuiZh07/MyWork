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
	app.Post("/register", controller.HandleRegistration)
	app.Post("/login", controller.HandleLogin)



	// Start server
	log.Fatal(app.Listen(":8081"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
