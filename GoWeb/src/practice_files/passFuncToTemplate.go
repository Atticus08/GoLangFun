package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// Pointer to template
var tpl *template.Template

// Engineer object
type engineer struct {
	Name string
}

// Functions being mapped to our template
var tplFuncs = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// Parse the template html file
func init() {
	/**
	When passing functions we can't do the normal way of parsing the template filename,
	and assigning a pointer to the template. We need to first pass in the functions to
	the file we're parsing. So, we do that by first creating a NEW template, tell it
	the functions we need to use, and then parse the static template file, letting it
	know the functions that need to be used for the new template.
	*/
	fmt.Println("I'm in init tpl = ", tpl)
	tpl = template.Must(template.New("").Funcs(tplFuncs).ParseFiles("tplFunc.gohtml"))
}

func main() {
	fmt.Println("I'm in main tpl = ", tpl)
	// Create index.html file to push template over to.
	newFile, error := os.Create("index.html")
	if error != nil {
		log.Println("There was an error creating the file", error)
	}
	defer newFile.Close()

	// Write out to our index.html file, and pass data to our template
	matt := engineer{"maTt schAEFer "}
	john := engineer{"John Wetula"}

	engineers := []engineer{matt, john}
	data := struct{ Engineer []engineer }{
		engineers,
	}

	error = tpl.ExecuteTemplate(newFile, "tplFunc.gohtml", data)
	if error != nil {
		log.Fatalln(error)
	}
}

// Pull out the first 3 characters of a string
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}
