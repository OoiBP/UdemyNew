package main

import "fmt"

const (
	zero  = 2019 + iota
	one   = 2019 + iota
	two   = 2019 + iota
	three = 2019 + iota
	four  = 2019 + iota
)

func main() {
	fmt.Println(zero, one, two, three, four)
}

/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/
