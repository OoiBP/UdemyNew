package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	x := customErr{
		info: "lots of confusion, and sadness",
	}

	foo(x)
	fmt.Printf("%T \n", x)
	s := x.Error()
	fmt.Printf("%T", s)
}

func foo(e error) {
	fmt.Println("From foo:")
	fmt.Println(e)
}

/*
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
*/
