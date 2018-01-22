package main

import "fmt"

// Person Object
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Agent Object
type Agent struct {
	Person
	lastName  string // Overriding in Go
	suitColor string
	codeName  string
}

// Get full name from person object
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Overriding the Person fullName function
func (a Agent) fullName() string {
	return a.firstName + " " + a.lastName
}

func structsPractice() {
	var tom = Agent{
		Person: Person{
			firstName: "Tom",
			lastName:  "Fritz",
			age:       30,
		},
		lastName:  "Fritzy",
		suitColor: "Black",
		codeName:  "Atticus08",
	}
	fmt.Println(tom.fullName())
	fmt.Println(tom.lastName)
}
