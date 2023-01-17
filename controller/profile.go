package controller

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	// "NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"log"
)

func LoadCreateNewProfile(c *fiber.Ctx) error {
	// Get social media platform name from the json file
	// return to the option in html file

	dataJson, err := ioutil.ReadFile("database/socialMedia.json")
	if err != nil {
		log.Fatal(err)
	}

	var data []model.SocialMedia
	err = json.Unmarshal(dataJson, &data)
	if err != nil {
		log.Fatal(err)
	}
	var mediaName []string
	for _, socialMedia := range data {
		mediaName = append(mediaName, socialMedia.PlatformName)
	}

	return c.Render("createProfile", mediaName)
}
