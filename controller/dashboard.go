package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
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
	userID := sess.Get(model.USER_ID)
	var userName string
	var userUniversity string

	// Get user name and university from database based on user's email
	err = database.DB.QueryRow("SELECT name, university FROM users WHERE email = $1", userEmail).Scan(&userName, &userUniversity)
	if err != nil {
		fmt.Print("Error when getting user name and university from database (dashboard.go)")
		log.Fatal(err)
	}

	// Check if the user has uploaded their own profile picture
	var profilePicture string
	_, err = os.Stat("avatar/" + userID.(string) + ".png")
	if err == nil {
		// If the user has uploaded their own profile picture, use it
		profilePicture = "avatar/" + userID.(string) + ".png"
	} else {
		// If the user hasn't uploaded their own profile picture, use the default one
		profilePicture = "avatar/user.png"
	}


	return c.Render("dashboard", fiber.Map{
		"ProfilePicture": profilePicture,
		"UserName":       userName,
		"UserUniversity": userUniversity,
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

