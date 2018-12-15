package main

import (
	"fmt"
)

func main() {
	for i :=1; i<=10; i++ {
		fmt.Println(i)
	}

	for i := 65; i <=90; i++ {
		fmt.Println(i)
		for j :=0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	bd := 1991
	for bd <= 2018 {
		fmt.Println(bd)
		bd++
	}

	year := 1991
	for {
		if year > 2019 {
			break
		}
		fmt.Println(year)
		year++
	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}


	x := "Rishabh"
	if x == "Rishabh" {
		fmt.Println(x)
	} else if x == "Bond" {
		fmt.Println("Hello", x)
	} else {
		fmt.Println("No option found")
	}

	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}


	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}