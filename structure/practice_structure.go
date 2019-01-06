package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	favFlavours []string
}

func main() {
	p1 := person {
		first: "Rishabh",
		last: "Rusia",
		favFlavours: []string{
			"Chocolate",
			"Rum and coke",
		},
	}

	p2 := person {
		first: "Prateek",
		last: "Rusia",
		favFlavours: []string{
			"Vanilla",
			"Strawberry",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println("\n")

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}
}