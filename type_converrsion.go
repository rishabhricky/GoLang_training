package main

import (
	"fmt"
)

type pizza int

var x pizza
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 45

	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", x)
}