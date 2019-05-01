package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		fav:   []string{"chocolate"},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		fav:   []string{"strawberry", "mint"},
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.fav {
		fmt.Println(i, v)
	}

	fmt.Println(p2)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.fav {
		fmt.Println(i, v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
