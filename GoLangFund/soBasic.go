package main

import (
	"fmt"
)

func soBasic() {
	number, isEven := half(900)
	fmt.Println("Half result: ", number)
	fmt.Println("The number is even: ", isEven)

	numbers := []int{10, 1000, 1, 999, 8, 76}
	fmt.Println("Here's the max number in our list of numbers: ", findMax(numbers...))

}

func half(number int) (int, bool) {
	return number / 2, number%2 == 0
}

func findMax(numbers ...int) int {
	var largest int
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}
