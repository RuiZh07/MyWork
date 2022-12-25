package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

type University struct {
	Name     string `json:"School Name"`
	Email    string `json:"URL"`
	City     string `json:"City"`
	Location string `json:"State"`
}

func main() {
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

	db, err = sql.Open("postgres", "postgres://cyw:cyw@localhost:5432/wacave?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

	defer db.Close()
	
	// Iterate over the slice of University structs and print the name, email, and location
	for _, university := range universities {

		_, err = db.Exec("INSERT INTO universities (name, domain, city, state) VALUES ($1, $2, $3, $4)",
			university.Name, university.Email, university.City, university.Location)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Name:", university.Name)
		// fmt.Println("Email:", university.Email)
		// fmt.Println("Location:", university.Location)
	}

	fmt.Print("done")
	
}
