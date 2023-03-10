package middleware

import (
	"NFC_Tag_UPoint/controller"
	"NFC_Tag_UPoint/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/gofiber/template/html"
	"log"
	"time"
	"os"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() {
	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		KeyGenerator:  utils.UUIDv4,
		CookieName:     "_csrf",
		CookieSameSite: "Strict",
		CookieSecure:   true,
		CookieHTTPOnly: true,

	}
	limiterConfig := limiter.Config{
		Max: 20,
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
	})

	// Setup middleware session
	log.Print("Setting up middleware session")

	// Routes
	app.Get("/", index)

	// This is Get request routes for user without authentication to view public profile
	app.Get("/page/:publicProfileLink", controller.LoadPublicProfile)
	app.Get("/page/avatar/:filename", controller.ServeAvatar)

	// This is Get request routes for user without authentication to access public tag
	app.Get("/tag/:tagHash", controller.LoadNFCPage)

	appPost := fiber.New(fiber.Config{
		Views: engine,
	})
	appPost.Use(limiter.New(limiterConfig),csrf.New(csrfConfig))

	// This is Post request routes for user without authentication to access public tag
	appPost.Post("/activateTag", controller.ActivateNFC)

	NoAuth := app.Group("/auth")
	NoAuth.Use(setAuth())

	// This is Get request routes for user without authentication

	NoAuth.Get("/SelectUniversity", controller.LoadUniversitySelection)
	NoAuth.Get("/login", controller.LoadLoginPage)

	//Setup NoAuthPost to limit the request reducing server load
	NoAuthPost := app.Group("/auth")
	NoAuthPost.Use(csrf.New(csrfConfig), limiter.New(limiterConfig))

	// This is Post request routes for user without authentication
	NoAuthPost.Post("/createAccount", controller.HandleUniversitySelection)
	NoAuthPost.Post("/register", controller.HandleRegistration)
	NoAuthPost.Post("/login", controller.HandleLogin)

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
	setting.Get("/avatar/:filename", controller.ServeAvatar)

	settingPost := admin.Group("/setting")
	settingPost.Use(limiter.New(limiterConfig), csrf.New(csrfConfig))

	settingPost.Post("/editInfo", controller.EditPersonalInfo)

	//Setup adminPost to limit the request reducing server load
	adminPost := app.Group("/user")
	adminPost.Use(limiter.New(limiterConfig),
		csrf.New(csrfConfig))

	admin.Post("/logout", controller.Logout)

	profilePost := adminPost.Group("/profile")
	profilePost.Use(csrf.New(csrfConfig))
	profilePost.Post("/createProfile", controller.CreateNewProfile)
	profilePost.Post("/deleteProfile", controller.DeleteProfile)
	profilePost.Post("/setAsPrimary", controller.SetAsPrimaryProfile)
	profilePost.Post("/setProfileLink", controller.CreateProfileLink)

	// Post for tag
	tagPost := adminPost.Group("/manageTag")
	tagPost.Use(limiter.New(limiterConfig), csrf.New(csrfConfig))

	tagPost.Post("/updateTagActivation", controller.DeactivateNFC)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}
