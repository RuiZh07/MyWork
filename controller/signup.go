package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"strings"
)

var universityData = make(map[string]model.University)

func HandleRegistration(c *fiber.Ctx) error {
	// Get the form values
	userName := template.HTMLEscapeString(c.FormValue("userName"))
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
		return UnexpectedError(c, err, "HandleRegisteration(signup.go)")
	}

	// Check if the email is already in use
	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return UnexpectedError(c, err, "HandleRegisteration(signup.go)")
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
		return UnexpectedError(c, err, "HandleRegisteration(signup.go)")
	}

	// Update user count in universities table
	_, err = database.DB.Exec("UPDATE universities SET user_numbers = user_numbers + 1 WHERE name = $1", university)
	if err != nil {
		return UnexpectedError(c, err, "HandleRegisteration(signup.go)")
	}

	c.Render("login", fiber.Map{
		"SuccessfullyRegistered": "Registered Successfully, Please Login With Your Account",
	})

	return nil
}

func LoadUniversitySelection(c *fiber.Ctx) error {

	dataJSON, err := ioutil.ReadFile("database/universityData.json")
	if err != nil {
		return UnexpectedError(c, err, "LoadUniversitySelection (signup.go)")
	}

	// Unmarshal the JSON data into a slice of UniversityData structs.
	var data []model.UniversityData
	err = json.Unmarshal(dataJSON, &data)
	if err != nil {
		return UnexpectedError(c, err, "LoadUniversitySelection (signup.go)")
	}
	var uName []string
	for _, university := range data {
		universityData[university.Name] = model.University{URL: university.Email, City: university.City, State: university.Location}
		uName = append(uName, university.Name)
	}

	return c.Render("universitySelectRegister", uName)

}

func HandleUniversitySelection(c *fiber.Ctx) error {
	universitySelected := c.FormValue("university")

	// Look up the university data in the universityData map.
	university, ok := universityData[universitySelected]
	if !ok {
		err := errors.New("university not found")
		return UnexpectedError(c, err, "HandleUniversitySelection (signup.go)")
	}

	return c.Render("signup", fiber.Map{
		"UniversityName":   universitySelected,
		"UniversityDomain": university.URL,
	})

}

func checkInputValidation(userName string, email string, password string, confirmPassword string, domain string) string {

	if userName == "" {
		return "Name can't be empty"
	}

	if password == "" {
		return "Password can't be empty"
	}

	if password != confirmPassword {
		return "Password mismatched"
	}

	if !strings.Contains(email, domain) {

		return "Please use your university email " + domain
	}

	//check if email is ending with .edu
	if !strings.HasSuffix(email, ".edu") {
		return "Please use your university email " + domain
	}
	return ""
}
