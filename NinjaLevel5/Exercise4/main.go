package main

import "fmt"

func main() {
	v := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	fmt.Println(v)
}

/*Create and use an anonymous struct.*/
