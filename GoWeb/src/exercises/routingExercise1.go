package exercises

import (
	"fmt"
	"net/http"
)

func rootFunc1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "You're at the root, bra")
}

func dogFunc1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "That be a dog")
}

func meFunc1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Tom Fritz")
}

// RoutingExercise1 - Using handlerfunc, wrtie a set of routes to the
// default ServeMux.
func RoutingExercise1() {
	http.HandleFunc("/", rootFunc1)
	http.HandleFunc("/dog/", dogFunc1)
	http.HandleFunc("/me/", meFunc1)

	// Use the default ServeMux by using "nil" as the handler
	// This way we don't have to create a new instance of ServeMux (if it's not needed)
	http.ListenAndServe(":8080", nil)
}
