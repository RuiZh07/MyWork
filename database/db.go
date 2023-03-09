package database

import (
	"NFC_Tag_UPoint/model"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

var DB *sql.DB

func Setup() {
	dsn := "postgres://admin:admin@localhost:5432/wacave?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	var count int

	// Create the users table if not exist.
	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'users'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Println("Creating users table")
		_, err = DB.Exec(`
			CREATE TABLE users (
				user_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				name text NOT NULL,
				email text NOT NULL,
				password text NOT NULL,
				university text NOT NULL,
				profilePicture text,
				profileLink text,
				role text NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the nfcTag table if not exist
	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'nfcTag'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Println("Creating nfcTag table")

		_, err = DB.Exec(`
			CREATE TABLE nfcTag (
				nfc_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				name text,
				tagHash VARCHAR(255),
				user_email text,
				activated BOOLEAN NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`)
	}

	// Create the universities table.
	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'universities'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Println("Creating university table")

		_, err = DB.Exec(`
			CREATE TABLE universities (
				university_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				domain VARCHAR(255) NOT NULL,
				city VARCHAR(255) NOT NULL,
				state VARCHAR(255) NOT NULL,
				user_numbers int NOT NULL
			);
		`)

		if err != nil {
			log.Fatal(err)
		}

		bytes, err := ioutil.ReadFile("database/universityData.json")
		if err != nil {
			log.Fatal(err)
		}

		// Parse the JSON data into a slice of UniversityData structs.
		var universities []model.UniversityData
		err = json.Unmarshal(bytes, &universities)
		if err != nil {
			log.Fatal(err)
		}

		//Loop through the universities and insert them into the database.
		for _, university := range universities {
			_, err = DB.Exec(`
				INSERT INTO universities (name, domain, city, state, user_numbers)
				VALUES ($1, $2, $3, $4, $5);
			`, university.Name, university.Email, university.City, university.Location, 0)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("Inserted ", len(universities), " universities")
	}

	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'profiles'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Println("Creating profile page table")

		_, err = DB.Exec(`
			CREATE TABLE profiles (
				profile_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				user_id INTEGER NOT NULL,
				user_email TEXT NOT NULL,
				name TEXT NOT NULL,
				activation BOOLEAN NOT NULL,
				link1 TEXT,
				link2 TEXT,
				link3 TEXT,
				link4 TEXT,
				link5 TEXT,
				link6 TEXT,
				link7 TEXT,
				link8 TEXT,
				link9 TEXT,
				link10 TEXT
			);
		`)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("All tables are created")
}
