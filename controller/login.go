package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func LoadLoginPage(c *fiber.Ctx) error {
	// Render login.html template
	return c.Render("login", nil)
}

// HandleLogin handles user login requests
func HandleLogin(c *fiber.Ctx) error {
	// Get the form values
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Query the user's record from the database
	var hashedPassword string
	var userID int
	err := database.DB.QueryRow("SELECT password, user_id FROM users WHERE email = $1", email).Scan(&hashedPassword, &userID)
	if err == sql.ErrNoRows {
		// No user with that email was found
		return c.Render("login", fiber.Map{"ErrorMessage": "No account associated with email: " + email})
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

	// The email and password are correct, log the user in
	sess, sessErr := model.Store.Get(c)
	if sessErr != nil {
		log.Fatal("Error when getting session info")
	}

	userIDStr := fmt.Sprintf("%d", userID)
	sess.Set(model.USER_ID, userIDStr)
	sess.Set(model.AUTH_KEY, true)
	sess.Set(model.USER_EMAIL, email)

	sessErr = sess.Save()
	if sessErr != nil {
		log.Fatal("Error when saving session info")
	}

	// The email and password are correct, log the user in
	return c.Redirect("/user/dashboard")
}
