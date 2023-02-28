package database

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type NFCTag struct {
	NFChash    string `json:"NFChash"`
	NFCName    string `json:"nfcName"`
	Activation bool   `json:"activation"`
}

func generateHash() string {
	b := make([]byte, 32)
	rand.Read(b)
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateNFC() {
	// Define the filename for the JSON file.
	filename := "database/nfctag.json"

	// Create an empty JSON file if it doesn't exist.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	}

	// Read the existing JSON data from the file.
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Decode the JSON data into a slice of NFCTag structs.
	var nfcTags []NFCTag
	if len(jsonFile) > 0 {
		err = json.Unmarshal(jsonFile, &nfcTags)
		if err != nil {
			panic(err)
		}
	}

	numTags := 100

	// Generate the specified number of NFC tags.
	for i := 0; i < numTags; i++ {
		// Generate a unique hash value.
		var nfcHash string
		var isUnique bool

		for !isUnique {
			nfcHash = generateHash()

			// Check if the hash already exists in the slice of NFCTag structs.
			isUnique = true
			for _, tag := range nfcTags {
				if tag.NFChash == nfcHash {
					isUnique = false
					break
				}
			}
		}

		// Create a new NFC tag.
		nfcTag := NFCTag{
			NFChash:    nfcHash,
			NFCName:    nfcHash[:8],
			Activation: false,
		}

		// Append the new NFC tag to the slice.
		nfcTags = append(nfcTags, nfcTag)
	}

	// Encode the updated slice of NFCTag structs to JSON format.
	jsonData, err := json.MarshalIndent(nfcTags, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the JSON data to the file.
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func CheckNfcAmount() {
	filename := "database/nfctag.json"

	// Read the existing JSON data from the file.
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Decode the JSON data into a slice of NFCTag structs.
	var nfcTags []NFCTag
	if len(jsonFile) > 0 {
		err = json.Unmarshal(jsonFile, &nfcTags)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No NFC generated")
		return
	}

	fmt.Println("total NFC generated:", len(nfcTags))
}
