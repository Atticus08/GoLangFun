package servingFiles

import (
	"fmt"
	"net/http"
	"os"
)

// ServingWithHTTP - Handle routing for Boomer
func ServingWithHTTP() {
	http.HandleFunc("/", root)
	http.HandleFunc("/boomer", boomerServeContent)
	http.ListenAndServe(":8080", nil)
}

// Handles root endpoint for API
func rootFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(res, "You're at the root of the API. Don't you want to see Boomer?")
}

func boomerServeContent(res http.ResponseWriter, req *http.Request) {
	// Open an image from our images folder
	file, error := os.Open("../img/boomer.jpg")
	if error != nil {
		http.Error(res, "File Not Found", 404)
		return
	}
	defer file.Close()

	fileStats, error := file.Stat()
	if error != nil {
		http.Error(res, "File Not Found", 404)
		return
	}

	// Serve content is used to tell when the file was last modified, and this is mainly for
	// HTTPs ETag, which is used for web cache validation.
	http.ServeContent(res, req, file.Name(), fileStats.ModTime(), file)
}
