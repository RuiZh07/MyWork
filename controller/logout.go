package controller

import (
	"NFC_Tag_UPoint/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info")
	}

	sess.Destroy()

	if sess.Get(model.AUTH_KEY) == nil {
		log.Println("empty AUTH_KEY")
	}
	
	log.Print("User log out")
	return c.Redirect("/")
}
