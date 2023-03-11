package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"os"
)

func LoadDeleteAccount(c *fiber.Ctx) error {
	return c.Render("deleteAccount", nil)
}

func HandleDeleteAccount(c *fiber.Ctx) error {
	email := c.FormValue("email")

	// Query user's university from the database
	var university string
	err := database.DB.QueryRow("SELECT university FROM users WHERE email = $1", email).Scan(&university)
	if err == sql.ErrNoRows {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}

	// Minus 1 from the number of users in the university
	_, err = database.DB.Exec("UPDATE universities SET user_numbers = user_numbers - 1 WHERE name = $1", university)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}

	// Check if the user has an nfc tag, if no, skip this step
	var nfcTagCount int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM nfcTag WHERE user_email = $1", email).Scan(&nfcTagCount)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}
	// If the user has an nfc tag, deactivate the nfc tag by setting the email to null,
	//	 change activated to false remove created_at and set name as first 5 character of tagHash
	if nfcTagCount > 0 {
		_, err = database.DB.Exec("UPDATE nfcTag SET user_email = null, activated = false, created_at = null, name = substring(tagHash from 1 for 5) WHERE email = $1", email)
		if err != nil {
			return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
		}
	}

	// Delete profiles that are associated with the user
	_, err = database.DB.Exec("DELETE FROM profiles WHERE user_email = $1", email)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}

	// Delete profile pictures that are associated with the user from the file system under avatar folder
	// Get the profile picture name from the database
	var profilePicture string
	err = database.DB.QueryRow("SELECT profilePicture FROM users WHERE email = $1", email).Scan(&profilePicture)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}
	// Delete the profile picture from the file system
	if profilePicture != "user.png" {
		err = os.Remove("avatar/" + profilePicture)
		if err != nil {
			return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
		}
	}

	// Delete the user's record from the database
	_, err = database.DB.Exec("DELETE FROM users WHERE email = $1", email)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}

	// Delete the user's session
	sess, err := model.Store.Get(c)
	if err != nil {
		return UnexpectedError(c, err, "HandleDeleteAccount (deleteAccount.go)")
	}
	sess.Destroy()

	return c.Redirect("/")
}
