package controller

import (
	"NFC_Tag_UPoint/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoadDashboard(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info in dashboard")
	}
	userEmail := sess.Get(model.USER_EMAIL)

	return c.Render("dashboard", fiber.Map{
		"UserEmail": userEmail,
	})
}

func LoadProfilePage(c *fiber.Ctx) error {
	return nil
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