package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// User - The Project's JSON Model
type User struct {
	Name             string    `json:"name"`
	Age              int       `json:"age"`
	DataJoined       time.Time `json:"date_joined"`
	AwesomenessLevel int       `json:"awesomeness_level"`
}

// Users - Slice of the User JSON model
type Users []User

func main() {
	// Create new router, and define the trailing slash behavior for
	// new routes. In this instance, we don't care if it's /path/ or /path
	router := mux.NewRouter().StrictSlash(true)

	// Set up the routes for the router to handle
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/user", UserIndexHandler)
	router.HandleFunc("/user/{userId}", UserShowIndexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// IndexHandler for when endpoint is called.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// UserIndexHandler is the handler for the /user route.
func UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{"Tom", 30, time.Now(), 10},
		User{"Erin", 28, time.Now().Add(10 * time.Second), 10},
	}
	json.NewEncoder(w).Encode(users)
}

// UserShowIndexHandler is the handler for the /user/{userId} route.
// For right now, we're only pulling the user id from the route.
func UserShowIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	fmt.Fprintln(w, "The User's Id: ", userID)
}
