package middleware

import (
	"NFC_Tag_UPoint/controller"
	"NFC_Tag_UPoint/model"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
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
	model.Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24,
	})

	// Setup middleware session
	log.Print("Setting up middleware session")

	// Routes
	app.Get("/", index)

	//Todo
	// Setup "/page/:profileLink" route
	app.Get("/page/:profileLink", controller.LoadPublicProfile)
	app.Get("/page/avatar/:filename", controller.ServeAvatar)


	NoAuth := app.Group("/auth")
	NoAuth.Use(setAuth())

	// This is Get request routes for user without authentication

	NoAuth.Get("/signup", controller.LoadRegister)
	NoAuth.Get("/login", controller.LoadLoginPage)

	//Setup NoAuthPost to limit the request reducing server load
	NoAuthPost := app.Group("/auth")
	NoAuthPost.Use(limiter.New())
	// This is Post request routes for user without authentication
	NoAuthPost.Post("/selectU", controller.HandleUniversitySelection)
	NoAuthPost.Post("/register", controller.HandleRegistration)
	NoAuthPost.Post("/login", controller.HandleLogin)

	admin := app.Group("/user")
	admin.Use(checkAuth())
	admin.Get("/dashboard", controller.LoadDashboard)

	// Todo
	// Complete each of the get request setup
	admin.Get("/profilePage", controller.LoadProfilePage)
	admin.Get("/createProfileLink", controller.LoadCreateNewProfileLink)
	admin.Get("/manageTag", controller.ManageTag)
	admin.Get("/requestTag", controller.RequestTag)
	admin.Get("/setting", controller.LoadSettingPage)

	// Avatar
	admin.Get("/avatar/:filename", controller.ServeAvatar)

	// Profile
	profile := admin.Group("/profile")
	profile.Get("/createNewProfile", controller.LoadCreateNewProfile)
	profile.Get("/:id", controller.DisplayProfile)

	// Setting
	setting := admin.Group("/setting")
	setting.Get("/editInfo", controller.LoadEditInfo)
	setting.Get("/avatar/:filename", controller.ServeAvatar)

	setting.Post("/editInfo", controller.EditPersonalInfo)

	//Setup adminPost to limit the request reducing server load
	adminPost := app.Group("/user")
	adminPost.Use(limiter.New(limiter.Config{
		Max: 20,
	}))

	admin.Post("/logout", controller.Logout)

	profilePost := adminPost.Group("/profile")
	profilePost.Post("/createProfile", controller.CreateNewProfile)
	profilePost.Post("/deleteProfile", controller.DeleteProfile)
	profilePost.Post("/setProfileLink", controller.CreateProfileLink)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
