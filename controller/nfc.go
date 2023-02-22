package controller

import (
	"NFC_Tag_UPoint/database"
	//"NFC_Tag_UPoint/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

/*
This function is used to load the NFC page
activation: boolean
  - true: NFC tag is activated
  - false: NFC tag is not activated

Return:
  - error: error message
*/
func LoadNFCPage(c *fiber.Ctx) error {
	// Get the activation status of the NFC tag
	activate := checkActivate(c)

	// Render the NFC page if the NFC tag is activated
	if activate {
		reditrecActivatedTag(c)
	}

	// Render the NFC page if the NFC tag is not activated
	return c.Render("activateTag", nil)
}

/*
This function is to verify if the NFC tag is activated from the database
Parameters:
  - c: fiber context

Return:
  - activated: boolean
*/
func checkActivate(c *fiber.Ctx) bool {
	// Get the tag hash from the URL
	tagHash := c.Params("tagHash")
	activated := false

	// Check if the tag is activated
	err := database.DB.QueryRow("SELECT activated FROM nfcTag WHERE tagHash = $1", tagHash).Scan(&activated)
	if err != nil {
		fmt.Print("Error when getting tag activation status from database (nfc.go)")
		log.Fatal(err)
	}

	return activated
}

/*
This function is to get the profile page of the NFC tag
Parameters:
  - c: fiber context
  - tagHash: string
  - activated: boolean
  - profile: model.Profile
  - profileInfo: model.ProfileData

Return:
  - profileInfo: model.ProfileData
*/
func reditrecActivatedTag(c *fiber.Ctx) error {
	// Get the tag hash from the URL
	tagHash := c.Params("tagHash")

	// Get the user email of the tag
	var userEmail string
	err := database.DB.QueryRow("SELECT user_email FROM nfcTag WHERE tagHash = $1", tagHash).Scan(&userEmail)
	if err != nil {
		fmt.Print("Error when getting user email from database (nfc.go)")
		log.Fatal(err)
	}

	// Get public profile's profile_link from database and redirect to the public profile page
	var profileLink string
	err = database.DB.QueryRow("SELECT profile_link FROM profiles WHERE user_email = $1 and activation = $2", userEmail, true).Scan(&profileLink)
	if err != nil {
		fmt.Print("Error when getting profile link from database (nfc.go)")
		log.Fatal(err)
	}

	return c.Redirect("/profile/" + profileLink)
}

/*
This function is to activate the NFC tag
Parameters:
  - c: fiber context
  - tagHash: string

Return:
  - error: error message
*/
func ActivateNFC(c *fiber.Ctx) error {
	// Get the tag hash from the URL
	tagHash := c.Params("tagHash")
	userEmail := c.FormValue("userEmail")
	confirmEmail := c.FormValue("confirmEmail")

	if userEmail != confirmEmail {
		return c.Render("activateTag", fiber.Map{
			"ErrorMessage": "Emails do not match",
		})
	}

	// check if the user email is in the database
	_, err := database.DB.Exec("SELECT user_email FROM users WHERE user_email = $1", userEmail)
	if err != nil {
		return c.Render("activateTag", fiber.Map{
			"ErrorMessage": "There is no account associated with this email, please check your email or create an account",
		})
	}

	// Activate the NFC tag
	_, err = database.DB.Exec("UPDATE nfcTag SET activated = $1 and user_email = $2 WHERE tagHash = $3", true, userEmail, tagHash)
	if err != nil {
		fmt.Print("Error when activating NFC tag (nfc.go)")
		log.Fatal(err)
	}

	// Redirect to the NFC page
	return c.Redirect("/nfc/" + tagHash)
}
