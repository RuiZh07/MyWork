package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

type University struct {
	Name     string `json:"School Name"`
	Email    string `json:"URL"`
	City     string `json:"City"`
	Location string `json:"State"`
}

func LoadUniversityData() {
	// Read the JSON file into a byte slice
	bytes, err := ioutil.ReadFile("universityData.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of University structs
	var universities []University
	err = json.Unmarshal(bytes, &universities)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", "postgres://admin:admin@localhost:5432/wacave?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Iterate over the slice of University structs and print the name, email, and location
	for _, university := range universities {

		// Make sure you create `universities` table in database before you import
		//
		// CREATE TABLE universities (
		// 	name VARCHAR(255) NOT NULL,
		// 	domain VARCHAR(255) NOT NULL,
		// 	city VARCHAR(255) NOT NULL,
		// 	state VARCHAR(255) NOT NULL
		//   );

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM universities WHERE domain = $1", university.Email).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		if count > 0 {
			// University is already in database
			continue
		} else{
			_, err = db.Exec("INSERT INTO universities (name, domain, city, state) VALUES ($1, $2, $3, $4)",
				university.Name, university.Email, university.City, university.Location)
			if err != nil {
				log.Fatal(err)
		}
		}

	}

	fmt.Println("done")

}
