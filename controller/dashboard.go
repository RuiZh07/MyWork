package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strings"
)

func LoadDashboard(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info in dashboard")
	}

	userEmail := sess.Get(model.USER_EMAIL)
	var userName string
	var userUniversity string
	var profilePicture string
	var profileLink sql.NullString

	// Get user name and university from database based on user's email
	err = database.DB.QueryRow("SELECT name, university, COALESCE(profilePicture, ''), profileLink FROM users WHERE email = $1", userEmail).Scan(&userName, &userUniversity, &profilePicture, &profileLink)
	if err != nil {
		fmt.Print("Error when getting user name and university from database (dashboard.go)")
		log.Fatal(err)
	}

	if !profileLink.Valid {
		profileLink.String = "Profile Link Not Set Yet, Go to Profile Page to Set"
	} else {
		// If the user has set their profile link, add the domain name to the link
		domain := c.BaseURL()
		profileLinkWithDomain := domain + "/page/" + profileLink.String
		profileLink.String = profileLinkWithDomain
	}

	// Check if the user has uploaded their own profile picture

	_, err = os.Stat("avatar/" + profilePicture)
	if err == nil {
		// If the user has uploaded their own profile picture, use it
		profilePicture = "avatar/" + profilePicture
	} else {
		// If the user hasn't uploaded their own profile picture, use the default one
		profilePicture = "avatar/user.png"
	}

	return c.Render("dashboard", fiber.Map{
		"ProfilePicture": profilePicture,
		"UserName":       userName,
		"UserUniversity": userUniversity,
		"ProfileLink":    profileLink.String,
	})
}

func ServeAvatar(c *fiber.Ctx) error {
	//Get URL path
	path := c.Path()
	// Splitting URL with "/"
	segments := strings.Split(path, "/")
	// Get the last segment in URL
	avatarPath := segments[len(segments)-1]

	return c.SendFile("avatar/" + avatarPath)
}
func ManageTag(c *fiber.Ctx) error {
	return nil
}

func RequestTag(c *fiber.Ctx) error {
	return nil
}
