package main

import (
	"fmt"
	"strconv"
)

func main() {
	for index := 0; index < 5; index++ {
		fmt.Println("Print " + strconv.Itoa(index))
	}
	fmt.Println()

	i := 0
	for i < 5 {
		fmt.Println("Print " + strconv.Itoa(i))
		i++
	}

	/*
		! range iterates over elements in a variety of data structures.
	*/
	nums := []int{5, 7, 8, 9, 1, 2}
	for index, num := range nums {
		fmt.Println("Print " + strconv.Itoa(index) + " | Number: " + strconv.Itoa(num))
	}
}
