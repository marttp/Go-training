package main

import "fmt"

func main() {
	// sliceExample := make([]float64, 5)
	// fmt.Println(sliceExample)

	// arrayExample := []float64{1, 2, 0}
	// sliceExample := append(arrayExample, 4, 5)
	// fmt.Println(sliceExample)

	arrayExample := []float64{1, 2, 0}
	sliceExample := make([]float64, 2)
	copy(sliceExample, arrayExample)
	fmt.Println(sliceExample)
}
