package main

import "fmt"

func main() {
	year := 1992
	for {
		if year <= 2019 {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}
}

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/
