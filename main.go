package main

import (
	"NFC_Tag_UPoint/database"
	"NFC_Tag_UPoint/middleware"
	// "time"
)

func main() {

	// // log.Println("Waitting for server to boot in 10s")
	// time.Sleep(10 * time.Second)

	// Start dabase
	database.Setup()

	// Create table in database
	// database.CreateTable()

	

	

	// Start web server
	middleware.Setup()
}
