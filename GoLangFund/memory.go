package main

import "fmt"

const metersToYards float64 = 1.09377

func memory() {
	var meters float64

	fmt.Println("Enter in the meters you swam: ")

	// Scan standard input, and we need to look at the actual memory location to
	// grab inputted value.
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards. ")
}
