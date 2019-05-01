package main

import "fmt"

func main() {
	list := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	list["lapsap_pui"] = []string{`Glass`, `Plastic`, `Tin`}

	for i, v := range list {
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}
}

/*
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
*/
