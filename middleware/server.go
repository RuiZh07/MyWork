package middleware

import (
	"NFC_Tag_UPoint/controller"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"log"
	"time"
)

var (
	store      *session.Store
	AUTH_KEY   string = "authenticated"
	USER_EMAIL string = "user_email"
)

func Setup() {

	// Initialize standard go html template engine
	engine := html.New("templates", ".html")

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup session cookie
	log.Print("Setting up session cookie")
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24,
	})

	// Setup middleware session
	log.Print("Setting up middleware session")

	// Routes
	app.Get("/", index)
	app.Get("/*", routeNotExist)

	NoAuth := app.Group("/auth")
	NoAuth.Use(setAuth())

	NoAuth.Get("/signup", controller.LoadRegister)
	NoAuth.Get("/login", controller.LoadLoginPage)
	NoAuth.Post("/selectU", controller.HandleUniversitySelection)
	NoAuth.Post("/register", controller.HandleRegistration)
	NoAuth.Post("/login", HandleLogin)

	admin := app.Group("/user")
	admin.Use(checkAuth())
	admin.Get("/dashboard", controller.LoadDashboard)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
