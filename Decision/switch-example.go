package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("Enter your grade : ")
	var inputGrade string
	fmt.Scan(&inputGrade)
	// fmt.Scanf("%d", &inputScore)

	// ! Another way to make input and parse to int
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// inputScore, _ := strconv.Atoi(scanner.Text())
	// fmt.Println(inputScore)
	switch strings.ToLower(inputGrade) {
	case "a":
		fmt.Print("Score range 80-100")
	case "b":
		fmt.Print("Score range 70-79")
	case "c":
		fmt.Print("Score range 60-69")
	case "d":
		fmt.Print("Score range 50-59")
	default:
		fmt.Print("Default Print")
	}

}
