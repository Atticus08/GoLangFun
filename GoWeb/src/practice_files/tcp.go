package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func tcp() {
	//Over TCP, listen for all available interfaces on port 8080
	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		log.Panic(error)
	}
	defer listener.Close()

	// Loop through until we see a connection
	for {
		connection, error := listener.Accept()
		if error != nil {
			log.Println(error)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	// Kill the connnection after 5 seconds
	timeOut := connection.SetDeadline(time.Now().Add(5 * time.Second))
	if timeOut != nil {
		log.Println("The Connection Timed Out!")
	}
	// Set up a scanner to read from the TCP connection
	scanner := bufio.NewScanner(connection)

	// Loop until scanner has scanned through everything, and stopped.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer connection.Close()

	fmt.Println("Connection Closed")
}
