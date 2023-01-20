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


func CreateNewProfile(c *fiber.Ctx) error{

	var mediaPlatform [] string
	var mediaAccountID [] string
	var mediaLink [] string
	itemIndex := 1
	idIndex := 1

	// get the input item
	for i := 0; i <= 10; i++{
		item := fmt.Sprintf("platform-%d", itemIndex)
		if c.FormValue(item) != ""{
			mediaPlatform = append(mediaPlatform, c.FormValue(item))
		}
		id := fmt.Sprintf("mediaID-%d", idIndex)
		if c.FormValue(id) != ""{
			mediaAccountID = append(mediaAccountID, c.FormValue(id))
		}
		itemIndex++
		idIndex++
	}
	

	fmt.Print(mediaPlatform)
	fmt.Print(mediaAccountID)

	profileName := c.FormValue("profileName")

	dataJson, err := ioutil.ReadFile("database/platformLinks.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON into a map of media platforms to URLs
	var mediaURLs map[string]string
	json.Unmarshal(dataJson, &mediaURLs)
	fmt.Println(mediaURLs)

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)

	
	// ISSUE, NOT IMPORTING TO MEDIALINK
	for index, account := range mediaAccountID{
		fmt.Println(mediaPlatform[index])
		if url, ok := mediaURLs[mediaPlatform[index]]; ok{
			link := url + account
			mediaLink = append(mediaLink, link)
		}
	}
	
	//Fix the insert statement to avoid duplicate the row,
	// TODO: update the json file for social media into only 1
	for index, link := range mediaLink{
		column := fmt.Sprintf("link%d", index+1)
		_, err = database.DB.Exec((fmt.Sprintf("INSERT INTO profiles (user_id, user_email, name, activation, %s) VALUES ($1, $2, $3, $4, $5)", column)),
			userID, userEmail, profileName, true, link)
		
		if err != nil{
			log.Fatal(err)
		}
	}

	return c.Redirect("/user/profilePage")

}

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