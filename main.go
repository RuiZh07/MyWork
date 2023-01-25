package main

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/middleware"
	// "time"
)

func main() {

	// // log.Println("Waitting for server to boot in 10s")
	// time.Sleep(10 * time.Second)

	// Create table in database
	// database.CreateTable()

	// Load University Data into Database
	// database.LoadUniversityData()

	// Start dabase
	database.Setup()

	// Start web server
	middleware.Setup()
}
