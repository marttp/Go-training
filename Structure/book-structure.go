package main

import (
	"fmt"
)

type Person struct {
	name  string
	money int
	age   int
}

func main() {

	var mart Person
	mart.name = "Mart"
	mart.money = 100
	mart.age = 22
	fmt.Println(mart.name)

	// ! Create short line
	martExample := Person{name: "Mart", money: 100, age: 22}
	fmt.Println(martExample.name)
}
