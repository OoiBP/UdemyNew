package main

import "fmt"

func main() {
	x := 42
	y := 24

	a := x == y
	b := x <= y
	c := x >= y
	d := x != y
	e := x < y
	f := x > y

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

/*
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
*/
