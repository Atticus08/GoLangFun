package servingFiles

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ServingWithIO - Handle routing for Boomer
func ServingWithIO() {
	http.HandleFunc("/", root)
	http.HandleFunc("/boomer", showMeBoomer)
	http.ListenAndServe(":8080", nil)
}

// Handles root endpoint for API
func root(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(res, "You're at the root of the API. Don't you want to see Boomer?")
}

// Retrieves the picture of Boomer from our images folder
func showMeBoomer(res http.ResponseWriter, req *http.Request) {
	// Open an image from our images folder
	file, error := os.Open("../img/boomer.jpg")
	if error != nil {
		http.Error(res, "File Not Found", 404)
		return
	}
	defer file.Close()

	// Copy the file data to our response writer
	io.Copy(res, file)
}
