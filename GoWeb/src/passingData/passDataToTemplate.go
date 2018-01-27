package main

import (
	"log"
	"os"
	"text/template"
)

type artist struct {
	name string
	age  int
}

type band struct {
	name   string
	rating int
}

type boxOfStuff struct {
	Artists []artist
	Bands   []band
}

func passDataExample() {
	tpl := parseFile("tpl.gohtml")

	// Create index.html file to push template over to.
	newFile, error := os.Create("index.html")
	if error != nil {
		log.Println("There was an error creating the file", error)
	}
	defer newFile.Close()

	// Write out to our index.html file, and pass data to our template
	var travis = artist{"Travis Barker", 40}
	var mark = artist{"Mark Hoppus", 42}
	var tom = artist{"Tom Delonge", 41}
	var blink = band{"Blink-182", 5}
	var sum41 = band{"Sum 41", 5}

	names := []artist{travis, tom, mark}
	bands := []band{blink, sum41}
	stuff := boxOfStuff{
		Artists: names,
		Bands:   bands,
	}
	error = tpl.Execute(newFile, stuff)
	if error != nil {
		log.Fatalln(error)
	}
}

// Parse the template html file, and pass it to our tpl pointer
// If the file doesn't exist, "Must" panics, goes through the rest of the
// call stack until all function in the current goroutine have returned, and
// then the program will crash.
func parseFile(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}
