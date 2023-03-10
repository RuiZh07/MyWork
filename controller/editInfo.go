package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
	"golang.org/x/crypto/bcrypt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"html/template"
)

// Render setting.html template
func LoadSettingPage(c *fiber.Ctx) error {
	return c.Render("setting", nil)
}

// This function is to load edit personal info page from GET request
func LoadEditInfo(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		UnexpectedError(c, err, "LoadEditInfo(editInfo.go)")
	}

	userID := sess.Get(model.USER_ID)
	var userName string
	var userUniversity string
	var profilePicture string

	// Get user's name, university, profile picture from the database
	err = database.DB.QueryRow("SELECT name, university, profilePicture FROM users WHERE user_id = $1", userID).Scan(&userName, &userUniversity, &profilePicture)
	if err != nil {
		log.Print("Error from db, LoadEditInfo")
		UnexpectedError(c, err, "LoadEditInfo(editInfo.go)")
	}
	// Access to user's profile picture in the filesystem
	profilePicture = "avatar/" + profilePicture

	return c.Render("editPersonalInfo", fiber.Map{
		"ProfilePicture": profilePicture,
		"University":     userUniversity,
		"userName":       userName,
	})
}

// This function is to update user's information from POST request
func EditPersonalInfo(c *fiber.Ctx) error {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		UnexpectedError(c, err, "EditPersonalInfo(editInfo.go)")
	}

	userID := sess.Get(model.USER_ID)
	var userName string
	var userUniversity string
	var profilePicture string

	// Get user's name, university, profile picture from the database
	err = database.DB.QueryRow("SELECT name, university, profilePicture FROM users WHERE user_id = $1", userID).Scan(&userName, &userUniversity, &profilePicture)
	if err != nil {
		log.Print("Error from db, LoadEditInfo")
		UnexpectedError(c, err, "EditPersonalInfo(editInfo.go)")
	}
	// Access to user's profile picture in the filesystem
	profilePicture = "avatar/" + profilePicture

	// Update the new profile Picture
	fileFormateError := changeImage(c)
	if fileFormateError != "" {
		return c.Render("editPersonalInfo", fiber.Map{
			"FileFormateErrorMessage": fileFormateError,
			"ProfilePicture":          profilePicture,
			"University":              userUniversity,
			"userName":                userName,
		})
	}

	// Update user name
	userNameError := changeUsername(c)
	if userNameError != "" {
		return c.Render("editPersonalInfo", fiber.Map{
			"ErrorMessage":   userNameError,
			"ProfilePicture": profilePicture,
			"University":     userUniversity,
			"userName":       userName,
		})
	}

	// Update Password
	passwordError := changePassword(c)
	if passwordError != "" {
		return c.Render("editPersonalInfo", fiber.Map{
			"ErrorMessage":   passwordError,
			"ProfilePicture": profilePicture,
			"University":     userUniversity,
			"userName":       userName,
		})
	}

	return c.Redirect("/user/setting")
}

// This function allow user to upload new profile picture,
// If the file isn't supported, return an error message
// else, save the image and update user's profile picture
func changeImage(c *fiber.Ctx) string {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		UnexpectedError(c, err, "ChangeImage(editInfo.go)")
	}

	// Get userId, userEmail from session
	userID := sess.Get(model.USER_ID)
	userEmail := sess.Get(model.USER_EMAIL)

	file, err := c.FormFile("newImage")

	// check If user doesn't upload new image
	if file == nil {
		return ""
	}
	if err != nil {
		UnexpectedError(c, err, "ChangeImage(editInfo.go)")
	}

	src, err := file.Open()
	if err != nil {
		log.Print(src)
	}
	defer src.Close()

	var profilePicture string
	err = database.DB.QueryRow("SELECT profilePicture FROM users WHERE email = $1", userEmail).Scan(&profilePicture)
	if err != nil {
		log.Print("Error when getting user name and university from database (editInfo.go, updateImage)")
		UnexpectedError(c, err, "ChangeImage(editInfo.go)")
	}

	// Check if the format is supported
	format := file.Header.Get("Content-Type")
	if format != "image/jpeg" && format != "image/png" && format != "image/jpg" {
		return "Not Supported File Formate"
	}

	// Decode the image
	img, format, err := image.Decode(src)
	if err != nil {
		UnexpectedError(c, err, "ChangeImage(editInfo.go)")
	}

	// Compress the image
	resizedImg := resize.Resize(100, 0, img, resize.Lanczos3)

	var fileName string

	// Set the file format based on the uploaded file
	switch format {
	case "jpeg":
		fileName = userID.(string) + ".jpeg"
		break
	case "jpg":
		fileName = userID.(string) + ".jpg"
		break
	case "png":
		fileName = userID.(string) + ".png"
		break
	default:
		fileName = userID.(string) + ".jpeg"
		break
	}

	// Remove current profile picture from filesystem
	if profilePicture != "user.png" {
		err = os.Remove("avatar/" + profilePicture)
		if err != nil {
			UnexpectedError(c, err, "ChangeImage (editInfo.go)")
		}
	}

	// Create the destination file
	dstFile, err := os.Create("avatar/" + fileName)
	if err != nil {
		UnexpectedError(c, err, "ChangeImage (editInfo.go)")
	}
	defer dstFile.Close()

	// Update the new image file to the users table
	_, err = database.DB.Exec("UPDATE users SET profilePicture = $1 WHERE email = $2", fileName, userEmail)
	if err != nil {
		log.Print("Error when getting user name and university from database (editInfo.go, updateImage)")
	}

	// Encode the image
	switch format {
	case "jpeg":
		if err = jpeg.Encode(dstFile, resizedImg, nil); err != nil {
			UnexpectedError(c, err, "ChangeImage (editInfo.go)")
		}
		break
	case "jpg":
		if err = jpeg.Encode(dstFile, resizedImg, nil); err != nil {
			UnexpectedError(c, err, "ChangeImage (editInfo.go)")
		}
		break
	case "png":
		if err = png.Encode(dstFile, resizedImg); err != nil {
			UnexpectedError(c, err, "ChangeImage (editInfo.go)")
		}
		break
	default:
		if err = jpeg.Encode(dstFile, resizedImg, nil); err != nil {
			UnexpectedError(c, err, "ChangeImage (editInfo.go)")
		}
		break
	}

	return ""
}

func changeUsername(c *fiber.Ctx) string {
	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		UnexpectedError(c, err, "ChangeUsername (editInfo.go)")
	}
	userID := sess.Get(model.USER_ID)

	userName := template.HTMLEscapeString(c.FormValue("userName"))
	_, err = database.DB.Exec("UPDATE users SET name = $1 WHERE user_id = $2", userName, userID)
	if err != nil {
		log.Print("Error when changing user name, changeUsername()")
		UnexpectedError(c, err, "ChangeUsername (editInfo.go)")
	}

	return ""

}

func changePassword(c *fiber.Ctx) string {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Print("Error when getting session info in dashboard")
		UnexpectedError(c, err, "ChangePassword(editInfo.go)")
	}

	userID := sess.Get(model.USER_ID)
	newPassword := c.FormValue("newPassword")
	confirmPassword := c.FormValue("confirmPassword")

	// Check if user wants to update password
	// If not, return nothing
	if newPassword == "" || confirmPassword == "" {
		return ""
	}

	if newPassword == "" {
		return "Password can't be empty"
	}

	if newPassword != confirmPassword {
		return "Password mismatched"
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Error when hasing password")
		UnexpectedError(c, err, "ChangePassword(editInfo.go)")
	}

	_, err = database.DB.Exec("UPDATE users SET password = $1 WHERE user_id = $2", hashedPassword, userID)
	if err != nil {
		log.Print("Error when changing user name, changeUsername()")
		UnexpectedError(c, err, "ChangePassword(editInfo.go)")
	}

	return ""
}
