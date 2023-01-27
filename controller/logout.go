package controller

import (
	"NFC_Tag_UPoint/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Logout(c *fiber.Ctx) error {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info")
	}

	sess.Destroy()

	return c.Redirect("/")
}
