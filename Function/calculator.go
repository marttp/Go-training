package main

import (
	"fmt"
)

func main() {

	var op1, op2 float32
	fmt.Scan(&op1, &op2)
	for index := 0; index < 4; index++ {
		switch index {
		case 0:
			fmt.Printf("Add = %f\n", add(op1, op2))
		case 1:
			fmt.Printf("Subtract = %f\n", subtract(op1, op2))
		case 2:
			fmt.Printf("Multiply = %f\n", multiply(op1, op2))
		case 3:
			fmt.Printf("Divide = %f\n", divide(op1, op2))
		}
	}
}

func add(op1 float32, op2 float32) float32 {
	result := (op1 + op2)
	return result
}

func subtract(op1 float32, op2 float32) float32 {
	result := (op1 - op2)
	return result
}

func multiply(op1 float32, op2 float32) float32 {
	result := (op1 * op2)
	return result
}

func divide(op1 float32, op2 float32) float32 {
	result := (op1 / op2)
	return result
}
