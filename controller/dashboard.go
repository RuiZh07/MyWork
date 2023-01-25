package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func LoadDashboard(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info in dashboard")
	}

	userEmail := sess.Get(model.USER_EMAIL)
	var userName string
	var userUniversity string

	// Get user name and university from database based on user's email
	err = database.DB.QueryRow("SELECT name, university FROM users WHERE email = $1", userEmail).Scan(&userName, &userUniversity)
	if err != nil {
		fmt.Print("Error when getting user name and university from database (dashboard.go)")
		log.Fatal(err)
	}

	return c.Render("dashboard", fiber.Map{
		"UserName":       userName,
		"UserUniversity": userUniversity,
	})
}

func ManageTag(c *fiber.Ctx) error {
	return nil
}

func RequestTag(c *fiber.Ctx) error {
	return nil
}

func UserSetting(c *fiber.Ctx) error {
	return nil
}
