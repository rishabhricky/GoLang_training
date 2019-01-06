package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Ricky")
	s1 := woo("Rishabh")
	fmt.Println(s1)
	x, y := mouse("Winnsers", "Outliers")
	fmt.Println(x)
	fmt.Println(y)

	sum := varaidicP(1,2,3,4,5,6,7,8)
	fmt.Println(sum)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello from bar to", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

// Variadic paramater, converts them in slice of int
func varaidicP(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	//var sum int
	for _, v := range x{
		sum += v 
	}
	// fmt.Println(sum)
	return sum
}