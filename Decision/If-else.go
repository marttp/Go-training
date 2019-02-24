package main

import (
	"fmt"
)

func main() {
	fmt.Print("Input number = ")
	var input float32
	fmt.Scan(&input)

	if input > 50 {
		fmt.Print("More than 50")
	} else if input < 50 {
		fmt.Print("Lower than 50")
	} else {
		fmt.Print("Equals 50")
	}

}
