package middleware

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

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
	sess, sessErr := model.Store.Get(c)
	if sessErr != nil {
		log.Fatal("Error when getting session info")
	}

	sess.Set(model.AUTH_KEY, true)
	sess.Set(model.USER_EMAIL, email)

	log.Print("Session email")
	log.Print(sess.Get(model.USER_EMAIL))
	sessErr = sess.Save()
	if sessErr != nil {
		log.Fatal("Error when saving session info")
	}

	// The email and password are correct, log the user in
	return c.Redirect("/user/dashboard")
}
