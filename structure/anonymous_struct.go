package main

import (
	"fmt"
)

func main() {

	p1 := struct {
		name   string
		age    int
		number int
	}{
		name:   "Rishabh",
		age:    27,
		number: 8015046211,
	}

	fmt.Println(p1)

	// More complex anonymous structure

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Rishabh",
		friends: map[string]int{
			"Ashish": 1,
			"Rohit":  2,
			"Amit":   3,
			"Karan":  4,
		},
		favDrinks: []string{
			"White Rum", "Beer", "Absolute Vodka",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}

}
