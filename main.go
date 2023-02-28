package main

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/middleware"
)

func main() {

	// Generate NFC tags
	//database.GenerateNFC()

	// Start dabase
	database.Setup()

	// Start web server
	middleware.Setup()
}
