package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"strings"
)

func HandleRegistration(c *fiber.Ctx) error {
	// Get the form values
	userName := c.FormValue("userName")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")
	university := c.FormValue("UniversityName")
	domain := c.FormValue("UniversityDomain")
	role := c.FormValue("role")
	defultProfilePicture := "user.png"

	// Check register input if its valid
	errorMessage := checkInputValidation(userName, email, password, confirmPassword, domain)
	if errorMessage != "" {
		return c.Render("signup", fiber.Map{
			"Name":             userName,
			"UniversityName":   university,
			"Email":            email,
			"UniversityDomain": domain,
			"ErrorMessage":     errorMessage,
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
	_, err = database.DB.Exec("INSERT INTO users (name, email, password, university, profilePicture, role, created_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)", userName, email, hashedPassword, university, defultProfilePicture, role)
	if err != nil {
		return err
	}

	// Update user count in universities table
	_, err = database.DB.Exec("UPDATE universities SET user_numbers = user_numbers + 1 WHERE name = $1", university)
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
	
	// Get university info from universities table in database
	var domain string
	err := database.DB.QueryRow("SELECT domain FROM universities WHERE name = $1", universitySelected).Scan(&domain)
	if err != nil {
		log.Print(err)
		log.Fatal("Error when getting university info from database")
	}

	return c.Render("signup", fiber.Map{
		"UniversityName":   universitySelected,
		"UniversityDomain": domain,
	})

}

func checkInputValidation(userName string, email string, password string, confirmPassword string, domain string) string {

	if userName == "" {
		return "Name can't be empty"
	}

	if password == "" {
		return"Password can't be empty"
	}

	if password != confirmPassword {
		return "Password mismatched"
	}

	if !strings.Contains(email, domain) {

		return "Please use your university email " + domain
	}

	return ""
}
