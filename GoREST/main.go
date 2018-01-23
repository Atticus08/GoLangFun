package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	fmt.Fprintf(w, "The User Index!")
}

// UserShowIndexHandler is the handler for the /user/{userId} route.
// For right now, we're only pulling the user id from the route.
func UserShowIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	fmt.Fprintln(w, "The User's Id: ", userID)
}
