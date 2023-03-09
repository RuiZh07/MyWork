package controller

import (
	"NFC_Tag_UPoint/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Logout(c *fiber.Ctx) error {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info")
		UnexpectedError(c, err, "logout.go")
	}

	sess.Destroy()

	return c.Redirect("/")
}
