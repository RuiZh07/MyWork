package controller

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/model"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
)

// Render setting.html template
func LoadSettingPage(c *fiber.Ctx) error {
	// sess, err := model.Store.Get(c)
	// if err != nil {
	// 	log.Fatal("Error when getting session info in dashboard")
	// }

	// userEmail := sess.Get(model.USER_EMAIL)
	// userID := sess.Get(model.USER_ID)
	// var userName string

	// // Get user name and university from database based on user's email
	// err = database.DB.QueryRow("SELECT name FROM users WHERE email = $1", userEmail).Scan(&userName)
	// if err != nil {
	// 	log.Print("Error when getting user name and university from database (dashboard.go)")
	// 	log.Fatal(err)
	// }

	// // Check if the user has uploaded their own profile picture
	// var profilePicture string
	// _, err = os.Stat("avatar/" + userID.(string) + ".png")
	// if err == nil {
	// 	// If the user has uploaded their own profile picture, use it
	// 	profilePicture = "avatar/" + userID.(string) + ".png"
	// } else {
	// 	// If the user hasn't uploaded their own profile picture, use the default one
	// 	profilePicture = "avatar/user.png"
	// }

	// return c.Render("setting", fiber.Map{
	// 	"ProfilePicture": profilePicture,
	// 	"userName": userName,
	// })
	return c.Render("setting", nil)
}

func LoadChangePicture(c *fiber.Ctx) error {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info in dashboard")
	}

	userEmail := sess.Get(model.USER_EMAIL)
	userID := sess.Get(model.USER_ID)
	var userName string

	// Get user name and university from database based on user's email
	err = database.DB.QueryRow("SELECT name FROM users WHERE email = $1", userEmail).Scan(&userName)
	if err != nil {
		log.Print("Error when getting user name and university from database (dashboard.go)")
		log.Fatal(err)
	}

	// Check if the user has uploaded their own profile picture
	var profilePicture string
	_, err = os.Stat("avatar/" + userID.(string) + ".png")
	if err == nil {
		// If the user has uploaded their own profile picture, use it
		profilePicture = "avatar/" + userID.(string) + ".png"
	} else {
		// If the user hasn't uploaded their own profile picture, use the default one
		profilePicture = "avatar/user.png"
	}

	return c.Render("changePicture", fiber.Map{
		"ProfilePicture": profilePicture,
	})
}

func UpdateImage(c *fiber.Ctx) error {

	sess, err := model.Store.Get(c)
	if err != nil {
		log.Fatal("Error when getting session info in dashboard")
	}

	// Get userId, userEmail from session
	userID := sess.Get(model.USER_ID)
	userEmail := sess.Get(model.USER_EMAIL)

	file, err := c.FormFile("newImage")
	if err != nil {
		log.Print(err)
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
	}

	// Check if the format is supported
	format := file.Header.Get("Content-Type")
	if format != "image/jpeg" && format != "image/png" && format != "image/jpg" {
		return c.Render("changePicture", fiber.Map{
			"ProfilePicture": "avatar/" + profilePicture,
			"ErrorMessage":   "Not Supported File Type!",
		})
	}

	// Decode the image
	img, format, err := image.Decode(src)
	if err != nil {
		log.Print(err)
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
	}

	// Create the destination file
	dstFile, err := os.Create("avatar/" + fileName)
	if err != nil {
		log.Print(err)
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
			log.Print(err)
		}
		break
	case "jpg":
		if err = jpeg.Encode(dstFile, resizedImg, nil); err != nil {
			log.Print(err)
		}
		break
	case "png":
		if err = png.Encode(dstFile, resizedImg); err != nil {
			log.Print(err)
		}
		break
	default:
		if err = jpeg.Encode(dstFile, resizedImg, nil); err != nil {
			log.Print(err)
		}
		break
	}

	return c.Redirect("/user/setting")
}

func LoadChangeUsername(c *fiber.Ctx) error {
	return c.Render("changeUsername", nil)
}

func LoadChangePassword(c *fiber.Ctx) error {
	return c.Render("changePassword", nil)
}
