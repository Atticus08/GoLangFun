package exercises

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func rootFunc2(res http.ResponseWriter, req *http.Request) {
	pushToHTML("../public_html/index.html", "You're at the root, bra")
}

func dogFunc2(res http.ResponseWriter, req *http.Request) {
	pushToHTML("../public_html/index.html", "<b>That be a dog, bra</b>")
}

func meFunc2(res http.ResponseWriter, req *http.Request) {
	pushToHTML("../public_html/index.html", "That be me, Tom Fritz")
}

// RoutingExercise2 - Using handlerfunc, parse template and store data in it
func RoutingExercise2() {
	// Set up routes
	http.HandleFunc("/", rootFunc2)
	http.HandleFunc("/dog/", dogFunc2)
	http.HandleFunc("/me", meFunc2)

	http.ListenAndServe(":8080", nil)
}

func pushToHTML(htmlFile string, data template.HTML) {
	// Get a pointer to point at our template file
	tpl := parseFile("../templates/tplData.gohtml")

	// Create html file from the template that was parsed
	newFile, error := os.Create(htmlFile)
	if error != nil {
		log.Println("Error creating HTML file", error)
	}
	defer newFile.Close()

	// Include our data in the parsed template, and spit the result out to
	// our HTML file.
	error = tpl.Execute(newFile, data)
	if error != nil {
		log.Fatalln(error)
	}
}

func parseFile(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}
