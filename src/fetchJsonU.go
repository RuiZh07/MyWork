package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	
)

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
	
	// Iterate over the slice of University structs and print the name, email, and location
	for _, university := range universities {
		fmt.Println("Name:", university.Name)
		fmt.Println("Email:", university.Email)
		fmt.Println("Location:", university.Location)
	}
	
}
