package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// This function is used to load the profile page
// if the user has a profile link then it will load the profile page
// otherwise it will redirect to the create profile link page
func LoadProfilePage(c *fiber.Ctx) error {

	// Get the user id from the session
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}
	var profileLink sql.NullString

	// Check if user already has user profile link
	err = database.DB.QueryRow("SELECT profileLink FROM users WHERE user_id = $1", sess.Get(model.USER_ID)).Scan(&profileLink)
	if err != nil {
		log.Fatal(err)
	}

	if !profileLink.Valid {
		return c.Redirect("/user/createProfileLink")
	}

	profile := model.ProfileMenu{
		ShowCreateProfileButton: canCreateNewProfile(c),
		ProfilePages:            profilePages(c),
	}
	return c.Render("profilePage", profile)
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

	// Get the media platform URL from the json file
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

	var mediaLink []string
	for index, account := range mediaAccountID {
		if url, ok := mediaURLs[mediaPlatform[index]]; ok {

			if !strings.Contains(account, ".com") {
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

func LoadCreateNewProfileLink(c *fiber.Ctx) error {
	return c.Render("createProfileLink", nil)
}

// create profile link for user's profile page
func CreateProfileLink(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)
	profileLink := c.FormValue("profileLink")

	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE profileLink = $1", profileLink).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return c.Render("createProfileLink", fiber.Map{
			"LinkMessage": "Cannot create link, link already exist",
		})
	}

	_, err = database.DB.Exec("UPDATE users SET profileLink = $1 WHERE user_id = $2 and email = $3", profileLink, userID, userEmail)
	if err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/user/profilePage")
}

// Load public profile page of user
func LoadPublicProfile(c *fiber.Ctx) error {

	//Get URL path
	path := c.Path()
	// Splitting URL with "/"
	segments := strings.Split(path, "/")
	// Get the last segment in URL
	profileLink := segments[len(segments)-1]

	var profile model.Profile
	var user model.User

	err := database.DB.QueryRow(`SELECT name, email, university, profilePicture, profileLink FROM users WHERE profileLink = $1`, profileLink).Scan(&user.Name, &user.Email, &user.University, &user.ProfilePicture, &user.ProfileLink)
	if err != nil {
		log.Print("Error when getting data from db (profile.go/LoadPublicProfile() ) ")
		log.Fatal(err)
	}

	err = database.DB.QueryRow(`SELECT * FROM profiles WHERE user_email = $1 and activation = $2 `, user.Email, true).Scan(&profile.ProfileID, &profile.UserID, &profile.UserEmail, &profile.Name, &profile.Activation, &profile.Link1, &profile.Link2, &profile.Link3,
		&profile.Link4, &profile.Link5, &profile.Link6, &profile.Link7, &profile.Link8, &profile.Link9, &profile.Link10)

	if err == sql.ErrNoRows {
		return c.Render("publicProfile", fiber.Map{
			"UserName":       user.Name,
			"ProfilePicture": user.ProfilePicture,
			"University":     user.University,
			"Error":          "No public profile set yet, please contact user to set a public profile",
		})
	} else if err != nil {
		log.Print("Error when getting profile from db (profile.go/LoadPublicProfile() ) ")
		log.Fatal(err)
	}

	var linkArray []string
	for i := 1; i <= 10; i++ {
		link := reflect.ValueOf(profile).FieldByName("Link" + strconv.Itoa(i))
		if link.FieldByName("Valid").Bool() && link.FieldByName("String").String() != "" {
			linkArray = append(linkArray, link.FieldByName("String").String())
		}
	}

	return c.Render("publicProfile", fiber.Map{
		"UserName":       user.Name,
		"ProfilePicture": user.ProfilePicture,
		"University":     user.University,
		"ProfileLinks":   linkArray,
	})
}
