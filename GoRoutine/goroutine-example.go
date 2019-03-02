package main

import (
	"fmt"
)

func main() {
	// firstExampleRoutineFunction(1)
	// ! ADD go keyword in front of function is working as routine
	go firstExampleRoutineFunction(1)
	var inputValue string
	fmt.Scanln(&inputValue)
}

func firstExampleRoutineFunction(value int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " : ", value)
	}
}

// Go routine = something like async function in Nodejs
// Working same time
