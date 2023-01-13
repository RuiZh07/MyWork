package controller

import (
	"NFC_Tag_UPoint/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
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

var domain string


func HandleRegistration(c *fiber.Ctx) error {
	// Get the form values
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")

	// Get university name
	var university string
	err := database.DB.QueryRow("SELECT name FROM universities WHERE domain = $1", domain).Scan(&university)
	if err != nil {
		return err
	}

	if checkInputValidation(email, password, confirmPassword) != "" {
		return c.Render("signup", fiber.Map{
			"UniversityName": university,
			"Email": email, 
			"UniversityDomain": domain,
			"ErrorMessage": checkInputValidation(email, password, confirmPassword),
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Check if the email is already in use
	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		// Email is already in use, return an error
		return c.Render("signup", fiber.Map{
			"UniversityName:":  university,
			"UniversityDomain": domain,
			"ErrorMessage":     "Email is already in use",
		})
	}

	// Insert the new user into the database
	_, err = database.DB.Exec("INSERT INTO users (email, password, university) VALUES ($1, $2, $3)", email, hashedPassword, university)
	if err != nil {
		return err
	}

	c.Render("login", fiber.Map{
		"SuccessfullyRegistered": "Registered Successfully, Please Login With Your Account",
	})

	return nil
}

func LoadRegister(c *fiber.Ctx) error {

	dataJSON, err := ioutil.ReadFile("database/universityData.json")
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
	err := database.DB.QueryRow("SELECT domain FROM universities WHERE name = $1", universitySelected).Scan(&domain)
	if err != nil {
		return err
	}

	return c.Render("signup", fiber.Map{
		"UniversityName":   universitySelected,
		"UniversityDomain": domain,
	})

}

func checkInputValidation(email string, password string, confirmPassword string) string{

	var errMessage string

	if password == "" {
		errMessage = "Password can't be empty"
		return errMessage
	}

	if password != confirmPassword {
		errMessage = "Password mismatched"
		return errMessage
	}

	if !strings.Contains(email, ".edu") {
		errMessage = "Invalid email"
		return errMessage
	}
	
	return errMessage
}