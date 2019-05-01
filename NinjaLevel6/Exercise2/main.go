package main

import "fmt"

func main() {
	arr := []int{32, 23, 42, 24}
	fmt.Println(foo(arr...))
	fmt.Println(bar(arr))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	return total
}

/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/
