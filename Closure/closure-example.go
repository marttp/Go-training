// Closure : create function by not define name of function
// Something like arrow function in Nodejs
package main

import (
	"fmt"
)

func main() {
	answer := func(op1, op2 int) int {
		return op1 + op2
	}
	fmt.Println(answer(5, 3))
}
