package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"strings"
)

var domain string

func HandleRegistration(c *fiber.Ctx) error {
	// Get the form values
	userName := c.FormValue("userName")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")

	// Get university name from /database/universityData.json
	var university string
	bytes, err := ioutil.ReadFile("database/universityData.json")
	if err != nil {
		log.Fatal(err)
	}

	var universities []model.UniversityData
	err = json.Unmarshal(bytes, &universities)
	if err != nil {
		fmt.Print("Error when loading university from json")
		log.Fatal(err)
	}
	for _, universityInfo := range universities {
		if universityInfo.Email == domain {
			university = universityInfo.Name
		}
	}

	// Check register input if its valid
	if checkInputValidation(userName, email, password, confirmPassword) != "" {
		return c.Render("signup", fiber.Map{
			"Name":             userName,
			"UniversityName":   university,
			"Email":            email,
			"UniversityDomain": domain,
			"ErrorMessage":     checkInputValidation(userName, email, password, confirmPassword),
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
			"Name":             userName,
			"UniversityName":   university,
			"Email":            email,
			"UniversityDomain": domain,
			"ErrorMessage":     "Email is already in use",
		})
	}

	// Insert the new user into the database
	_, err = database.DB.Exec("INSERT INTO users (name, email, password, university) VALUES ($1, $2, $3, $4)", userName, email, hashedPassword, university)
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
	var data []model.UniversityData
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
	bytes, err := ioutil.ReadFile("database/universityData.json")
	if err != nil {
		log.Fatal(err)
	}

	var universities []model.UniversityData
	err = json.Unmarshal(bytes, &universities)
	if err != nil {
		fmt.Print("Error when loading university from json")
		log.Fatal(err)
	}
	for _, universityInfo := range universities {
		if universityInfo.Name == universitySelected {
			domain = universityInfo.Email
		}
	}

	return c.Render("signup", fiber.Map{
		"UniversityName":   universitySelected,
		"UniversityDomain": domain,
	})

}

func checkInputValidation(userName string, email string, password string, confirmPassword string) string {

	var errMessage string

	if userName == "" {
		errMessage = "Name can't be empty"
		return errMessage
	}

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
