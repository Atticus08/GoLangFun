package main

import (
	"encoding/json"
	"fmt"
)

// Cyborg Object
// Utilizing unexported identifier, and json tags
type Cyborg struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	motto     string // Not exported, so JSON marshal will not see this
}

// Spy Object
type Spy struct {
	Cyborg    Cyborg
	SuitColor string `json:"suit_color"`
	CodeName  string `json:"-"`
}

// Song object, which is used to mess around with unmarshaling
type song struct {
	Artist string `json:"song_artist"`
	Title  string `json:"song_title"`
	Owner  string
}

func jsonMarshal() {
	fmt.Println("Marshalling...")
	// Marshal the heck out of that object!
	var p1 = Cyborg{"Tom", "Fritz", 31, "That's a terrible idea. What time?"}
	var tom = Spy{p1, "Black", "Atticus"}
	marshal(tom)

	fmt.Println("Unmarshalling...")
	// Unmarshal dat response bro
	var jsonResponse = []byte(`{"song_artist":"Blink-182", "song_title":"Feelin This", "song_owner":"Atticus"}`)
	var blinkSong song
	decodedResponse := unmarshal(jsonResponse, blinkSong)
	fmt.Println(decodedResponse)
	fmt.Printf("%T \n", decodedResponse)
}

func marshal(v interface{}) {
	byteStream, _ := json.Marshal(v)
	// This just prints out the string of characters of the structure
	// that was passed into Marshal.
	fmt.Println(byteStream)

	fmt.Printf("%T \n", byteStream)
	fmt.Println(string(byteStream))
}

func unmarshal(response []byte, v interface{}) interface{} {
	json.Unmarshal(response, &v)
	return v
}
