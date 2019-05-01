package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello, World")
}

func foo() {
	fmt.Println("From foo")
}

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
