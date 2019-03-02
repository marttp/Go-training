package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(3))
}

// recursive function
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
