package controller

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	// "NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"log"
	"fmt"
)

func LoadProfilePage(c *fiber.Ctx) error {
	profile := model.ProfileData{
		ShowCreateProfileButton: canCreateNewProfile(c),
		ProfilePages: profilePages(c),
	}
	return c.Render("profile", profile)
}


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

//	Todo
// func CreateNewProfile(c *fiber.Ctx) error{

// }

func canCreateNewProfile(c *fiber.Ctx) bool {

	var count int
	var userID int

	sess, err := model.Store.Get(c)
	if err != nil{
		fmt.Print("Error when getting session data (profile.go/canCreateNewProfile() ) ")
		log.Fatal(err)
	}
	userIDStr := sess.Get(model.USER_ID).(string)
	fmt.Sscanf(userIDStr, "%d", &userID)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM profiles WHERE user_id = $1", userID).Scan(&count)

	if err != nil{
		fmt.Print("Error when getting data from db (profile.go/canCreateNewProfile() ) ")
		log.Fatal(err)
	}

	if count < 3{
		return true
	}

	return false
}

func profilePages(c *fiber.Ctx) []string {

	var userID int
	sess, err := model.Store.Get(c)
	if err != nil {
		fmt.Print("Error when getting session data (profile.go/profilePage() ) ")
		log.Fatal(err)
	}

	userIDStr := sess.Get(model.USER_ID).(string)
	fmt.Sscanf(userIDStr, "%d", &userID)

	
	row, errs := database.DB.Query("SELECT name FROM profiles WHERE user_id = $1", userID)
	if errs != nil{
		fmt.Print("Error when getting data from db (profile.go/profilePage() ) ")
		log.Fatal(errs)
	}

	var profileNames [] string
	for row.Next() {
		var name string
		if err = row.Scan(&name); err != nil{
			log.Fatal(err)
		}
		profileNames = append(profileNames, name)
	}

	return profileNames
}