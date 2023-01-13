package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func setAuth() fiber.Handler {
	return checkAuth
}

func checkAuth(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		log.Fatal("Error when getting session info")
	}

	if sess.Get(AUTH_KEY) == nil {
		log.Print("No auth key in cookie")
	} else {
		c.Redirect("/dashboard")
	}

	return c.Next()
}
