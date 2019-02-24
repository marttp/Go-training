package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter text for Reader: ")
	text, _ := reader.ReadString('\n')

	fmt.Print("Enter text for Scanner: ")
	scanner.Scan()
	scanText := scanner.Text()

	fmt.Print("Reader : " + text)
	fmt.Print("Scanner : " + scanText)
}
