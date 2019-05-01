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
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/
