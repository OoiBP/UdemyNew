package main

import "fmt"

func main() {
	favSport := "Taekwon-Do"
	switch favSport {
	case "Badminton", "Basket Ball":
		fmt.Println("Lapsap")
	case "Taekwon-Do":
		fmt.Println("Good choice")
	}
}

/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/
