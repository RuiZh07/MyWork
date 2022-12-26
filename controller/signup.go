package controller

import (
	"database/sql"
	_ "github.com/lib/pq"
    "log"
    "golang.org/x/crypto/bcrypt"
    "github.com/gofiber/fiber/v2"
)


var db *sql.DB

func init() {
    var err error
    // Install postgresDB in your machine and change the `admin:admin` with your `username:password` and change `wacave` with your database name
    // make sure you create table in your database with following code
    // CREATE TABLE users (
    //     id serial PRIMARY KEY,
    //     email text NOT NULL,
    //     password text NOT NULL,
    //     university text NOT NULL
    // );
    
    db, err = sql.Open("postgres", "postgres://admin:admin@localhost:5432/wacave?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}


func HandleRegistration(c *fiber.Ctx) error {
    // Get the form values
    email := c.FormValue("email")
    password := c.FormValue("password")
    university := c.FormValue("university")

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Check if the email is already in use
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
    if err != nil {
        return err
    }
    if count > 0 {
        // Email is already in use, return an error
        return c.Status(400).SendString("Email is already in use")
    }

    // Insert the new user into the database
    _, err = db.Exec("INSERT INTO users (email, password, university) VALUES ($1, $2, $3)", email, hashedPassword, university)
    if err != nil {
        return err
    }

    c.Redirect("/login")
    return nil
}

func getUniversities(db *sql.DB) ([]string, error) {

	rows, err := db.Query("SELECT name FROM universities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var universities []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		universities = append(universities, name)
	}
	return universities, nil
}


func LoadRegister(c *fiber.Ctx) error{
	universities, err := getUniversities(db)
    
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error querying database")
	}
    
    return c.Render("signup", universities)
	
}