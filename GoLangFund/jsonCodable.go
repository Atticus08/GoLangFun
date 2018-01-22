package main

import (
	"encoding/json"
	"os"
)

// Dog object
type Dog struct {
	Name  string
	Breed string
	Age   int
}

func main() {
	var beren = Dog{
		Name:  "Beren",
		Breed: "Husky",
		Age:   3,
	}
	// Encode Beren object and display on terminal!
	json.NewEncoder(os.Stdout).Encode(beren)
}
