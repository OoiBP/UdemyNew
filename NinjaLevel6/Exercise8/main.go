package main

import "fmt"

func main() {
	bar := foo()
	bar()
}

func foo() func() {
	x := 0
	return func() {
		x++
		fmt.Println(x)
	}
}

/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/
