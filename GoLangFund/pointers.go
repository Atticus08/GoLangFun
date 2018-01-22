package main

import "fmt"

func pointers() {
	var a = 665
	var b = &a
	var c = *b

	fmt.Println("Memory address for a: ", &a)
	fmt.Println("Memory address b is pointing at: ", b) // Equal to memory address of "a"
	fmt.Println("Value in c: ", c)                      // Dereference b, and look at the value stored at the address

	var x = 10
	var y = &x // Pass x as reference to y

	zero(y)
	fmt.Println("Value in x: ", x) // I expect to see 0 in here
}

// Pass variable by reference, and set the value in that memory space to 0!
func zero(integer *int) {
	*integer = 0
}
