package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const input = "Be excellent to each other - Ted Logan"

func hashTable() {
	var hashTable = make([]int, 12)
	hashTable = scanHashInput(input, hashTable)
	fmt.Println(hashTable)
}

// Scan and hash the input.
func scanHashInput(input string, hashTable []int) []int {
	var scanner = bufio.NewScanner(strings.NewReader(input))

	// Split input based on space-separated words
	scanner.Split(bufio.ScanWords)

	// Go through scan and store in hash table.
	// Only storing the counts of each hash going into slice.
	for scanner.Scan() {
		n := hashWord(scanner.Text(), 12)
		fmt.Println("Word:", scanner.Text(), "Hashed to: ", n)
		fmt.Println("Hash Table at n: ", hashTable[n])
		hashTable[n]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading Input Error: ", err)
	}

	return hashTable
}

// Hash the word that was scanned from the Input
func hashWord(word string, numOfBuckets int) int {
	var sum int
	for _, val := range word {
		sum += int(val)
	}
	return sum % numOfBuckets
}
