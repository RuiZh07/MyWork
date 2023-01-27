package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	// "NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"encoding/json"
	"fmt"
	"log"
)

func LoadProfilePage(c *fiber.Ctx) error {
	profile := model.ProfileMenu{
		ShowCreateProfileButton: canCreateNewProfile(c),
		ProfilePages:            profilePages(c),
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

func CreateNewProfile(c *fiber.Ctx) error {

	var mediaPlatform []string
	var mediaAccountID []string
	var mediaLink []string
	itemIndex := 1
	idIndex := 1

	// get the input item
	for i := 0; i <= 10; i++ {
		item := fmt.Sprintf("platform-%d", itemIndex)
		if c.FormValue(item) != "" {
			mediaPlatform = append(mediaPlatform, c.FormValue(item))
		}
		id := fmt.Sprintf("mediaID-%d", idIndex)
		if c.FormValue(id) != "" {
			mediaAccountID = append(mediaAccountID, c.FormValue(id))
		}
		itemIndex++
		idIndex++
	}

	profileName := c.FormValue("profileName")

	dataJson, err := ioutil.ReadFile("database/platformLinks.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON into a map of media platforms to URLs
	var mediaURLs map[string]string
	json.Unmarshal(dataJson, &mediaURLs)

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)

	for index, account := range mediaAccountID {
		if url, ok := mediaURLs[mediaPlatform[index]]; ok {
			if !strings.Contains(".com", account) {
				link := url + account
				mediaLink = append(mediaLink, link)
			} else {
				mediaLink = append(mediaLink, account)
			}

		}
	}

	//Create the profile row if not exist, if user has same profile name then do nothing
	var profileID int
	err = database.DB.QueryRow("SELECT user_id FROM profiles WHERE user_id = $1 AND name = $2", userID, profileName).Scan(&profileID)
	switch {
	// If no row exist, then create new profile
	case err == sql.ErrNoRows:

		_, err = database.DB.Exec("INSERT INTO profiles (user_id, user_email, name, activation) VALUES ($1, $2, $3, $4)", userID, userEmail, profileName, false)

		if err != nil {
			log.Fatal(err)
		}

	case err != nil:
		log.Fatal(err)

	default:
		log.Print("Inserted new profile row into table")

	}

	// TODO: update the json file for social media into only 1
	for index, link := range mediaLink {
		column := fmt.Sprintf("link%d", index+1)
		_, err = database.DB.Exec((fmt.Sprintf("UPDATE profiles SET %s = $1 WHERE user_id = $2 AND name = $3", column)), link, userID, profileName)

		if err != nil {
			log.Fatal(err)
		}
	}

	return c.Redirect("/user/profilePage")

}

func canCreateNewProfile(c *fiber.Ctx) bool {

	var count int

	sess, err := model.Store.Get(c)
	if err != nil {
		fmt.Print("Error when getting session data (profile.go/canCreateNewProfile() ) ")
		log.Fatal(err)
	}
	userID := sess.Get(model.USER_ID)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM profiles WHERE user_id = $1", userID).Scan(&count)

	if err != nil {
		fmt.Print("Error when getting data from db (profile.go/canCreateNewProfile() ) ")
		log.Fatal(err)
	}

	if count < 3 {
		return true
	}

	return false
}

func profilePages(c *fiber.Ctx) []string {

	sess, err := model.Store.Get(c)
	if err != nil {
		fmt.Print("Error when getting session data (profile.go/profilePage() ) ")
		log.Fatal(err)
	}

	userID := sess.Get(model.USER_ID)

	row, errs := database.DB.Query("SELECT name FROM profiles WHERE user_id = $1", userID)
	if errs != nil {
		fmt.Print("Error when getting data from db (profile.go/profilePage() ) ")
		log.Fatal(errs)
	}

	var profileNames []string
	for row.Next() {
		var name string
		if err = row.Scan(&name); err != nil {
			log.Fatal(err)
		}
		profileNames = append(profileNames, name)
	}

	return profileNames
}

func DisplayProfile(c *fiber.Ctx) error {

	//Get URL path
	path := c.Path()
	// Splitting URL with "/"
	segments := strings.Split(path, "/")
	// Get the last segment in URL
	profileName := segments[len(segments)-1]

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)

	var profile model.Profile

	err = database.DB.QueryRow(`SELECT * FROM profiles WHERE user_id = $1 and user_email = $2 and name = $3`, userID, userEmail, profileName).Scan(&profile.ProfileID, &profile.UserID, &profile.UserEmail, &profile.Name, &profile.Activation, &profile.Link1, &profile.Link2, &profile.Link3,
		&profile.Link4, &profile.Link5, &profile.Link6, &profile.Link7, &profile.Link8, &profile.Link9, &profile.Link10)

	if err != nil {
		log.Fatal(err)
	}
	var linkArray []string
	for i := 1; i <= 10; i++ {
		link := reflect.ValueOf(profile).FieldByName("Link" + strconv.Itoa(i))
		if link.FieldByName("Valid").Bool() && link.FieldByName("String").String() != "" {
			linkArray = append(linkArray, link.FieldByName("String").String())
		}
	}

	profileInfo := model.ProfileData{
		ProfileName:  profileName,
		ProfileLinks: linkArray,
	}

	return c.Render("displayProfile", profileInfo)
}

func DeleteProfile(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)
	profileName := c.FormValue("profileName")

	_, err = database.DB.Exec("DELETE FROM profiles WHERE user_id = $1 and user_email = $2 and name = $3", userID, userEmail, profileName)
	if err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/user/profilePage")

}
