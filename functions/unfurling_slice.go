package main

import (
	"fmt"
)

func main() {
	xi := []int{1,2,3,4,5,6,7,8}
	fmt.Println(xi)
	
	// sum := varaidicP() - Will give zero, variadic means zero or more
	sum := varaidicP(xi...)
	fmt.Println(sum)

	unfurling("rishabh", "rishabh", "ricky")
}

func varaidicP(xi ...int) int {
	fmt.Println(xi)
	fmt.Printf("%T\n", xi)

	sum := 0
	//var sum int
	for _, v := range xi{
		sum += v 
	}
	// fmt.Println(sum)
	return sum
}

func unfurling(s string, xi ...string) {
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println(xi)
	fmt.Printf("%T\n", xi)

}