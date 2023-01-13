package middleware

import (
	"log"
	"strings"
	"github.com/gofiber/fiber/v2"
)

func setAuth() fiber.Handler {
	return noAuth
}

func noAuth(c *fiber.Ctx) error {
    sess, err := store.Get(c)

    if err != nil {
        log.Fatal("Error when getting session info")
    }

	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

    if sess.Get(AUTH_KEY) == nil {
        log.Print("No auth key in cookie")
    } 

    return c.Next()
}

func checkAuth() fiber.Handler {
	return auth
}

func auth(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		log.Fatal("Error when getting session info")
	}

	if sess.Get(AUTH_KEY) == nil {
		if strings.Split(c.Path(), "/")[1] != "auth" {
			c.Redirect("/auth/login")
		}
	}

	if strings.Split(c.Path(), "/")[1] == "user" {
		return c.Next()
	}

	return c.Next()
}