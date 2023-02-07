package main

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/middleware"
	// "time"
)

func main() {

	// Start dabase
	database.Setup()

	// Start web server
	middleware.Setup()
}
