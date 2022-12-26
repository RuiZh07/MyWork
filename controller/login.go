package controller

import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

// HandleLogin handles user login requests
func HandleLogin(c *fiber.Ctx) error {
    // Get the form values
    email := c.FormValue("email")
    password := c.FormValue("password")

    // Connect to the database
    db, err := sql.Open("postgres", "postgres://admin:admin@localhost/wacave?sslmode=disable")
    if err != nil {
        return err
    }
    defer db.Close()

    // Query the user's record from the database
    var hashedPassword string
    err = db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashedPassword)
    if err == sql.ErrNoRows {
        // No user with that email was found
        return c.Render("login", fiber.Map{"ErrorMessage": "Invalid email",})
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
    return c.SendString("Logged in successfully")
}


func LoadLoginPage(c *fiber.Ctx) error {
	// Render login.html template
	return c.Render("login", nil)
}
