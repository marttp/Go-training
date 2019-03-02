package main

import (
	"fmt"
)

func main() {
	fmt.Println(summation(5, 9, 8, 7, 6))
}

// ! variadic function - recieve argument
// ! by nobody know many of variable
func summation(args ...int) int {
	var total int

	// _ : means no define, start at position 0
	// num : means each number of argument
	for _, num := range args {
		total += num
	}
	return total
}
