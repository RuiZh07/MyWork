package model

// global class for university data
type UniversityData struct {
	Name     string `json:"School Name"`
	Email    string `json:"URL"`
	City     string `json:"City"`
	Location string `json:"State"`
}

type University struct {
	URL   string
	City  string
	State string
}
