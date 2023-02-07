package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

var DB *sql.DB

func Setup(){
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
		fmt.Println("Creating users table")
		_, err = DB.Exec(`
			CREATE TABLE users (
				user_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				name text NOT NULL,
				email text NOT NULL,
				password text NOT NULL,
				university text NOT NULL
			);
		`)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the universities table.
	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'universities'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Println("Creating university table")

		_, err = DB.Exec(`
			CREATE TABLE universities (
				university_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				domain VARCHAR(255) NOT NULL,
				city VARCHAR(255) NOT NULL,
				state VARCHAR(255) NOT NULL
			);
		`)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = DB.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'profiles'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Println("Creating profile page table")

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

	fmt.Println("All tables are created")
}