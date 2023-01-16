package controller

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	// "NFC_Tag_UPoint/database"
	"encoding/json"
	"NFC_Tag_UPoint/model"
	"log"
)

func LoadCreateNewProfile(c *fiber.Ctx) error {
	dataJson, err := ioutil.ReadFile("database/socialMedia.json")
	if err != nil {
		log.Fatal(err)
	}

	var data [] model.SocialMedia
	err = json.Unmarshal(dataJson, &data)
	if err != nil{
		log.Fatal(err)
	}
	var mediaName [] string
	for _, socialMedia := range data{
		mediaName = append(mediaName, socialMedia.PlatformName)
	}

	return c.Render("createProfile", mediaName)
}
