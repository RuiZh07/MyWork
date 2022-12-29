package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"strings"
)

type UniversityData struct {
	Name     string `json:"School Name"`
	Email    string `json:"URL"`
	City     string `json:"City"`
	Location string `json:"State"`
}

var db *sql.DB

var domain string

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

	// Get university name
	var university string
	err := db.QueryRow("SELECT name FROM universities WHERE domain = $1", domain).Scan(&university)
	if err != nil {
		return err
	}

	if !strings.Contains(email, ".edu") {

		return c.Render("signup", fiber.Map{
			"UniversityName": university,
			"UniversityDomain": domain,
			"ErrorMessage":     "Email Domain Not Supported",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// // Get university name
	// var university string
	// err = db.QueryRow("SELECT university FROM universities WHERE domain = $1", domain).Scan(&university)
	// if err != nil {
	// 	return err
	// }

	// Check if the email is already in use
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// Email is already in use, return an error
		return c.Render("signup", fiber.Map{
			"UniversityDomain": domain,
			"ErrorMessage":     "Email is already in use",
		})
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (email, password, university) VALUES ($1, $2, $3)", email, hashedPassword, university)
	if err != nil {
		return err
	}

	c.Redirect("/login")
	return nil
}

func LoadRegister(c *fiber.Ctx) error {

	dataJSON, err := ioutil.ReadFile("data/universityData.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of UniversityData structs.
	var data []UniversityData
	err = json.Unmarshal(dataJSON, &data)
	if err != nil {
		log.Fatal(err)
	}
	var uName []string
	for _, university := range data {
		uName = append(uName, university.Name)
	}

	return c.Render("universitySelectRegister", uName)

}

func HandleUniversitySelection(c *fiber.Ctx) error {
	universitySelected := c.FormValue("university")
	err := db.QueryRow("SELECT domain FROM universities WHERE name = $1", universitySelected).Scan(&domain)
	if err != nil {
		return err
	}

	return c.Render("signup", fiber.Map{
		"UniversityName": universitySelected,
		"UniversityDomain": domain,
	})

}