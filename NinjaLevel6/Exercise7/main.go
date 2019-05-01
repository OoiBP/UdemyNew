package main

import "fmt"

func main() {
	x := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	x()
	fmt.Println("")
	y := x
	y()
}

/*
Assign a func to a variable, then call that func
*/
