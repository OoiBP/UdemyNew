package main

import "fmt"

func main() {
	number := 42
	switch {
	case number == 42:
		fmt.Println("Number 42")
	case number == 43:
		fmt.Println("number 43")
	default:
		fmt.Println("Unknown")
	}
}

/*
Create a program that uses a switch statement with no switch expression specified.
*/
