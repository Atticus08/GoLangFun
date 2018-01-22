package main

import "fmt"

func slices() {
	var slice = make([]int, 0, 2)

	for i := 0; i < 8; i++ {
		slice = append(slice, i)
		fmt.Println("Len: ", len(slice), "Cap: ", cap(slice), "Slice: ", slice)
		fmt.Printf("Address: %p\n", slice)
	}
}
