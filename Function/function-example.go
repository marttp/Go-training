package main

import (
	"fmt"
)

func main() {
	introduce()
}

func introduce() {
	fmt.Println("================================")
	var name = "Thanaphoom Babparn"
	age := 22
	fmt.Println("Hello everyone.")
	fmt.Println("My name's ", name)
	fmt.Printf("I'm %v years old.\n", age)
	fmt.Println("================================")
}
