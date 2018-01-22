package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func blankID() {
	/**
	  Example of using a blank identifier. The GET HTTP command will return a
	  response, result, and an error. Instead of using result and error, we can
	  replace result and error with "_". Kind of like Swift.
	*/
	res, _ := http.Get("http://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
