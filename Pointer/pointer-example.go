package main

import (
	"fmt"
)

func main() {
	value := 15.95645
	// %f -> float number
	fmt.Printf("Value: %f \n", value)
	// %x -> address of variable
	fmt.Printf("Address: %x \n", &value)

	// ! *type -> create pointer
	var pointerExample *float64
	// Store address of variable
	pointerExample = &value
	fmt.Printf("Pointer direct to: %x \n", pointerExample)
	fmt.Printf("Address of Pointer: %x \n", &pointerExample)
	fmt.Printf("Value of Pointer: %f \n", *pointerExample)
	// fmt.Printf("Try this : %f \n", pointerExample)
	// fmt.Printf("Try this : %f \n", &pointerExample)

	// ! Manage data via pointer
	*pointerExample = 555.56465
	fmt.Printf("Value after changed: %f \n", value)
	fmt.Printf("Value of Pointer after changed: %f \n", *pointerExample)

}
