package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
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
	return c.Render("activateTag", fiber.Map{
		"tagHash": c.Params("tagHash"),
	})
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
		UnexpectedError(c, err, "checkActivate (nfc.go)")
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
		return UnexpectedError(c, err, "redirectActivatedTag (nfc.go)")
	}

	// Get public profile's profile_link from database and redirect to the public profile page
	var profileLink string
	err = database.DB.QueryRow("SELECT profileLink FROM users WHERE email = $1", userEmail).Scan(&profileLink)
	if err != nil {
		fmt.Print("Error when getting profile link from database (nfc.go)")
		return UnexpectedError(c, err, "redirectActivatedTag (nfc.go)")
	}

	return c.Redirect("/page/" + profileLink)
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
	tagHash := c.FormValue("tagHash")
	userEmail := c.FormValue("userEmail")
	confirmEmail := c.FormValue("confirmEmail")

	if userEmail != confirmEmail {
		return c.Render("activateTag", fiber.Map{
			"ErrorMessage": "Emails do not match",
		})
	}

	// check if the user email is in the database
	_, err := database.DB.Exec("SELECT email FROM users WHERE email = $1", userEmail)
	if err != nil {
		return c.Render("activateTag", fiber.Map{
			"ErrorMessage": "There is no account associated with this email, please check your email or create an account",
		})
	}

	// Activate the NFC tag
	_, err = database.DB.Exec("UPDATE nfcTag SET activated = $1, user_email = $2, created_at = COALESCE(created_at, CURRENT_TIMESTAMP) WHERE tagHash = $3", true, userEmail, tagHash)
	if err != nil {
		fmt.Print("Error when activating NFC tag (nfc.go)")
		return UnexpectedError(c, err, "ActivateNFC(nfc.go)")
	}

	// Redirect to the NFC page
	return c.Redirect("/tag/" + tagHash)
}

// Display the NFC tag associated with the user in settings page
func LoadNFCSetting(c *fiber.Ctx) error {
	// Get the user email from the session
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		return UnexpectedError(c, err, "LoadNFCSetting(nfc.go)")
	}

	userEmail := sess.Get(model.USER_EMAIL)

	// Create a slice to store the NFC tag name associated with the user from the database
	var nfcTags []model.NFCTag
	rows, err := database.DB.Query("SELECT nfc_id, name, activated, created_at FROM nfcTag WHERE user_email = $1", userEmail)
	if err != nil {
		fmt.Print("Error when getting NFC tag name from database (nfc.go)")
		return UnexpectedError(c, err, "LoadNFCSetting(nfc.go)")
	}
	// Store the NFC tag name in the slice
	for rows.Next() {
		var tempTag model.NFCTag
		err = rows.Scan(&tempTag.ID, &tempTag.Name, &tempTag.Activation, &tempTag.CreatedAt)
		if err != nil {
			fmt.Print("Error when scanning NFC tag name from database (nfc.go)")
			return UnexpectedError(c, err, "LoadNFCSetting(nfc.go)")
		}
		// Parse the timestamp string into a time.Time value
		timestamp, err := time.Parse(time.RFC3339Nano, tempTag.CreatedAt)
		if err != nil {
			fmt.Print("Error when parsing timestamp (nfc.go)")
			return UnexpectedError(c, err, "LoadNFCSetting(nfc.go)")
		}
		// Format the timestamp
		tempTag.CreatedAt = timestamp.Format("Jan 2006")
		nfcTags = append(nfcTags, tempTag)
	}

	// Render the NFC setting page
	return c.Render("NFCsetting", fiber.Map{
		"Tags": nfcTags,
	})

}

// Deactivate the NFC tag
func DeactivateNFC(c *fiber.Ctx) error {
	tagID := c.FormValue("tagID")
	activation := c.FormValue("tagActivation")

	// Get the user email from the session
	sess, err := model.Store.Get(c)
	if err != nil {
		return UnexpectedError(c, err, "DeactivateNFC (nfc.go)")
	}
	userEmail := sess.Get(model.USER_EMAIL)

	// Update the activation status of the NFC tag in the database
	_, err = database.DB.Exec("UPDATE nfcTag SET activated = $1 WHERE nfc_id = $2 AND user_email = $3", activation, tagID, userEmail)
	if err != nil {
		fmt.Print("Error when deactivating NFC tag (nfc.go)")
		return UnexpectedError(c, err, "DeactivateNFC (nfc.go)")
	}

	// Redirect to the NFC setting page
	return c.Redirect("/user/manageTag")
}
