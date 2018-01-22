package main

import "fmt"

func callbacks() {
	var numbers = []int{1, 5, 40, 20}
	var sum = 0
	visit(numbers, func(n int) {
		sum += n
	})
	fmt.Println("The sum is: ", sum)
}

func visit(numbers []int, callback func(n int)) {
	for _, n := range numbers {
		callback(n)
	}
}
