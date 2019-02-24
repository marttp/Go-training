package main

import "fmt"

/*
* Type of variable
- integer
- float
- boolean
- String
- Structure - Array, Map, Slice
* How to declare variable
- var money int = 5000
- money:=5000
:= ? find dynamical type
*/

func main() {
	// ! Example program - Introduce yourself with Varible
	var name = "Thanaphoom Babparn"
	age := 22
	fmt.Println("Hello everyone.")
	fmt.Println("My name's ", name)
	fmt.Printf("I'm %v years old.", age)
}
