package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

var db *sql.DB

type University struct {
	Name     string `json:"School Name"`
	Email    string `json:"URL"`
	City     string `json:"City"`
	Location string `json:"State"`
}

func init() {
	var err error
	// Install postgresDB in your machine and change the `admin:admin` with your `username:password` and change `wacave` with your database name
	// make sure you create table in your database with following code
	// CREATE TABLE users (
	//     id serial PRIMARY KEY,
	//     email text NOT NULL,
	//     password text NOT NULL,
	//     university text NOT NULL
	// );
	//
	// Make sure you create `universities` table in database before you import
	//
	// CREATE TABLE universities (
	// 	name VARCHAR(255) NOT NULL,
	// 	domain VARCHAR(255) NOT NULL,
	// 	city VARCHAR(255) NOT NULL,
	// 	state VARCHAR(255) NOT NULL
	//   );

	db, err = sql.Open("postgres", "postgres://admin:admin@localhost:5432/wacave?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func LoadUniversityData() {

	// Read the JSON file into a byte slice
	bytes, err := ioutil.ReadFile("database/universityData.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of University structs
	var universities []University
	err = json.Unmarshal(bytes, &universities)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the slice of University structs and print the name, email, and location
	for _, university := range universities {

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM universities WHERE domain = $1", university.Email).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		if count > 0 {
			// University is already in database
			continue
		} else {
			_, err = db.Exec("INSERT INTO universities (name, domain, city, state) VALUES ($1, $2, $3, $4)",
				university.Name, university.Email, university.City, university.Location)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

	fmt.Println("university data insertion done")

}

func CreateTable() {

	var count int

	// Create the users table if not exist.
	err := db.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'users'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Println("Creating users table")
		_, err = db.Exec(`
			CREATE TABLE users (
				id serial PRIMARY KEY,
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
	err = db.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'universities'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Println("Creating university table")

		_, err = db.Exec(`
			CREATE TABLE universities (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				domain VARCHAR(255) NOT NULL,
				city VARCHAR(255) NOT NULL,
				state VARCHAR(255) NOT NULL
			)
		`)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All tables are created")
}
