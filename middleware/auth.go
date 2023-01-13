package middleware

import (
	"log"
	"strings"
	"github.com/gofiber/fiber/v2"
)

func setAuth() fiber.Handler {
	return noAuth
}

// This is to check if user has session AUTH_KEY in their cookie, redirect
// user to /user/dashboard if they already signed in
func noAuth(c *fiber.Ctx) error {
    sess, err := store.Get(c)

    if err != nil {
        log.Fatal("Error when getting session info")
    }

	if strings.Split(c.Path(), "/")[1] == "auth" && sess.Get(AUTH_KEY) == nil{
		log.Print("No auth key in cookie")
		return c.Next()
	}

    return c.Redirect("/user/dashboard")
}

func checkAuth() fiber.Handler {
	return auth
}

// This is to check if user has session AUTH_KEY in their cookie, redirect
// user to /auth/login if they haven't signed in or session expired
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

	return nil
}