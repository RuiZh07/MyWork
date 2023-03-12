package middleware

import (
	"NFC_Tag_UPoint/controller"
	"NFC_Tag_UPoint/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"

	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func Setup() {

	limiterConfig := limiter.Config{
		Max: 200,
		Expiration: 1 * time.Minute,
	}
	file, err2 := os.OpenFile("database/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		log.Print(err2)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}

	// Initialize standard go html template engine
	engine := html.New("templates", ".html")

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Use middlewares for each route
	app.Use(
		logger.New(loggerConfig), // add Logger middleware with config

	)

	// Setup session cookie
	log.Print("Setting up session cookie")
	model.Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24,
		CookieSecure:  true,
	})

	csrfConfig := csrf.Config{
		// Next: func(c *fiber.Ctx) bool {
		// 	return true
		// },
		KeyLookup:         "cookie:_csrf",
		KeyGenerator:      utils.UUIDv4,
		CookieName:        "_csrf",
		CookieSameSite:    "Strict",
		CookieSecure:      true,
		CookieHTTPOnly:    true,
		CookieSessionOnly: true,
	}
	csrfProtection := csrf.New(csrfConfig)

	// Setup middleware session
	log.Print("Setting up middleware session")

	// Routes
	app.Get("/", index)

	// This is Get request routes for user without authentication to view public profile
	app.Get("/page/:publicProfileLink", controller.LoadPublicProfile)
	app.Get("/page/avatar/:filename", controller.ServeAvatar)

	// This is Get request routes for user without authentication to access public tag
	app.Get("/tag/:tagHash", controller.LoadNFCPage)

	appPost := app.Group("/")
	appPost.Use(limiter.New(limiterConfig))

	// This is Post request routes for user without authentication to access public tag
	appPost.Post("/activateTag", controller.ActivateNFC)

	NoAuth := app.Group("/auth")
	NoAuth.Use(setAuth())

	// This is Get request routes for user without authentication

	NoAuth.Get("/SelectUniversity", controller.LoadUniversitySelection)
	NoAuth.Get("/login", csrfProtection, controller.LoadLoginPage)

	//Setup NoAuthPost to limit the request reducing server load
	NoAuthPost := app.Group("/auth")
	NoAuthPost.Use(limiter.New(limiterConfig))

	// This is Post request routes for user without authentication
	NoAuthPost.Post("/createAccount", controller.HandleUniversitySelection)
	NoAuthPost.Post("/register", controller.HandleRegistration)
	NoAuthPost.Post("/login", csrfProtection, controller.HandleLogin)

	admin := app.Group("/user")
	admin.Use(checkAuth())
	admin.Get("/dashboard", controller.LoadDashboard)

	// Todo
	// Complete each of the get request setup
	admin.Get("/profilePage", controller.LoadProfilePage)
	admin.Get("/createProfileLink", controller.LoadCreateNewProfileLink)
	admin.Get("/manageTag", controller.LoadNFCSetting)
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
	setting.Get("/deleteAccount", controller.LoadDeleteAccount)
	setting.Get("/avatar/:filename", controller.ServeAvatar)

	settingPost := admin.Group("/setting")
	settingPost.Use(limiter.New(limiterConfig))

	settingPost.Post("/deleteAccount", csrfProtection, controller.HandleDeleteAccount)
	settingPost.Post("/editInfo", csrfProtection, controller.EditPersonalInfo)

	//Setup adminPost to limit the request reducing server load
	adminPost := app.Group("/user")
	adminPost.Use(limiter.New(limiterConfig))

	admin.Post("/logout", controller.Logout)

	profilePost := adminPost.Group("/profile")
	profilePost.Use(limiter.New(limiterConfig))
	profilePost.Post("/createProfile", csrfProtection, controller.CreateNewProfile)
	profilePost.Post("/deleteProfile", csrfProtection, controller.DeleteProfile)
	profilePost.Post("/setAsPrimary", csrfProtection, controller.SetAsPrimaryProfile)
	profilePost.Post("/setProfileLink", csrfProtection, controller.CreateProfileLink)

	// Post for tag
	tagPost := adminPost.Group("/manageTag")
	tagPost.Use(limiter.New(limiterConfig))

	tagPost.Post("/updateTagActivation", csrfProtection, controller.DeactivateNFC)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
