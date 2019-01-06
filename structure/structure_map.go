package main

import (
	"fmt"
)

// Level:5 q2 ??

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {
	p1 := person{
		first: "Rishabh",
		last:  "Rusia",
		favFlavours: []string{
			"Chocolate",
			"Rum and coke",
		},
	}

	p2 := person{
		first: "Prateek",
		last:  "Rusia",
		favFlavours: []string{
			"Vanilla",
			"Strawberry",
		},
	}
	// Store tye values in a map
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)
		}
		fmt.Println("--------------------")
	}
}
