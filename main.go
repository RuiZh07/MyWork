package main

import (
	"log"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
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
	app.Get("/signupPage", loadSignUp)
	app.Get("/loginPage", loadLogin)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}

func loadLogin(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func loadSignUp(c *fiber.Ctx) error {
	return c.Render("signup", nil)
}
