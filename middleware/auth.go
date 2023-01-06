package middleware

import (
	"NFC_Tag_UPoint/database"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

// NewMiddleware() to setup AuthMiddleware
func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "auth" {
		log.Print(c.Path())
		return c.Next()
	}

	if err != nil {
		log.Fatal("Err when getting authorized")
	}

	if sess.Get(AUTH_KEY) == nil {
		log.Print("No authorized auth_key")
		log.Print(sess.Get(AUTH_KEY))
		log.Print(sess.Get(USER_EMAIL))
	} else {
		c.Redirect("/auth/dashboard")
	}

	return c.Next()
}

// HandleLogin handles user login requests
func HandleLogin(c *fiber.Ctx) error {
	// Get the form values
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Query the user's record from the database
	var hashedPassword string
	err := database.DB.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		// No user with that email was found
		return c.Render("login", fiber.Map{"ErrorMessage": "Invalid email"})
	}
	if err != nil {
		return err
	}

	// Compare the provided password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// The password is incorrect
		return c.Render("login", fiber.Map{
			"ErrorMessage": "Invalid password",
		})
	}
	if err != nil {
		return err
	}

	log.Print("Get session info")
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		log.Fatal("Error when getting session info")
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_EMAIL, email)

	log.Print("Session email")
	log.Print(sess.Get(USER_EMAIL))
	sessErr = sess.Save()
	if sessErr != nil {
		log.Fatal("Error when saving session info")
	}

	// The email and password are correct, log the user in
	return c.Redirect("dashboard", 301)
}
