package main

import (
	"fmt"
)

func main() {
	// declare array
	// var arrayOfNumber[10] int
	arrayOfNumber := [10]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var total float64
	for _, num := range arrayOfNumber {
		total += num
	}
	fmt.Println(total / float64(len(arrayOfNumber)))
}
