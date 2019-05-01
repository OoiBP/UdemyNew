package main

import "fmt"

func main() {
	x := foo(func() int {
		value := 2 + 5
		return value
	})
	fmt.Println(x)
}

func foo(bar func() int) int {
	return bar() + 3
}

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/
