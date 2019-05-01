package main

import "fmt"

func main() {
	xstr := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(xstr)
	fmt.Printf("%T\n", xstr)

	for i := 0; i < len(xstr); i++ {
		for j := 0; j < len(xstr[i]); j++ {
			fmt.Printf("The index %d %d is %s\n", i, j, xstr[i][j])
		}
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
*/
