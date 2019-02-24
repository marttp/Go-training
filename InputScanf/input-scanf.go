package main

import (
	"fmt"
)

func main() {
	fmt.Print("Operand 1 = ")
	var operand1 float32
	// fmt.Scanf("%f", &operand1)
	fmt.Scan(&operand1)

	fmt.Print("Operand 2 = ")
	var operand2 float32
	// fmt.Scanf("%f", &operand2)
	fmt.Scan(&operand2)

	answer := operand1 + operand2
	fmt.Printf("Answer = %f \n", answer)
}
