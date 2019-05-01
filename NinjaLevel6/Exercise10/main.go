package main

import "fmt"

func main() {
	here := foo()
	here()
	here()
	here()
	another := foo()
	another()
	another()
	another()
	another()
}

func foo() func() {
	x := 0
	return func() {
		x++
		fmt.Println(x)
	}
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/
